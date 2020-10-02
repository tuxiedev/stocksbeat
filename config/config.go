// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config


type Config struct {
	Symbols 		[]string     `config:"symbols"`
	FinnhubToken	string		 `config:"finnhubToken"`
}

var DefaultConfig = Config{
	Symbols: make([]string, 0),
}
