package beater

// WebSocketResponseData structure reported from finnhub API documentation
type WebSocketResponseData struct {
	Price  float32 `json:"p"`
	Symbol string  `json:"s"`
	Time   int64   `json:"t"`
	Volume int32   `json:"v"`
}

// WebSocketResponse structure reported from finnhub API documentation
type WebSocketResponse struct {
	Data []WebSocketResponseData `json:"data"`
	Type string                  `json:"type"`
}
