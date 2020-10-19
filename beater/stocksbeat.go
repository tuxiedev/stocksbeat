package beater

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gorilla/websocket"

	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/logp"

	"github.com/tuxiedev/stocksbeat/config"
)

// stocksbeat configuration.
type stocksbeat struct {
	done   chan struct{}
	config config.Config
	client beat.Client
}

// New creates an instance of stocksbeat.
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	c := config.DefaultConfig
	if err := cfg.Unpack(&c); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	bt := &stocksbeat{
		done:   make(chan struct{}),
		config: c,
	}
	return bt, nil
}

// Run starts stocksbeat.
func (bt *stocksbeat) Run(b *beat.Beat) error {
	logp.Info("stocksbeat is running! Hit CTRL-C to stop it.")
	var err error
	bt.client, err = b.Publisher.Connect()
	if err != nil {
		return err
	}

	w, _, err := websocket.DefaultDialer.Dial("wss://ws.finnhub.io?token="+bt.config.FinnhubToken, nil)
	if err != nil {
		panic(err)
	}
	// logp.Info("Connected to finnhub websocket")

	for _, s := range bt.config.Symbols {
		msg, _ := json.Marshal(map[string]interface{}{"type": "subscribe", "symbol": s})
		w.WriteMessage(websocket.TextMessage, msg)
	}

	var msg WebSocketResponse
	for {
		select {
		case <-bt.done:
			w.Close()
			return nil
		default:
			err := w.ReadJSON(&msg)
			if err != nil {
				panic(err)
			}
			if msg.Type == "trade" {
				for _, trade := range msg.Data {
					bt.client.Publish(beat.Event{
						Timestamp: time.Unix(0, trade.Time*int64(time.Millisecond)),
						Fields: common.MapStr{
							"type": "trade",
							"symbol": trade.Symbol,
							"trade": common.MapStr{
								"price":  trade.Price,
								"volume": trade.Volume,
							},
						},
					})
				}
			}
		}
	}
}

// Stop stops stocksbeat.
func (bt *stocksbeat) Stop() {
	bt.client.Close()
	close(bt.done)
}
