package model

// Currency ...
type Currency struct {
	ID        int
	Symbol    string  `json:"symbol"`
	Price     float64 `json:"price_24h"`
	Volume    float64 `json:"volume_24h"`
	LastTrade float64 `json:"last_trade_price"`
}

type CurrencyInput struct {
	Symbol         string  `json:"symbol"`
	Price24H       float64 `json:"price_24h"`
	Volume24H      float64 `json:"volume_24h"`
	LastTradePrice float64 `json:"last_trade_price"`
}

type CurrencyOutput struct {
	Price     float64 `json:"price"`
	Volume    float64 `json:"volume"`
	LastTrade float64 `json:"last_trade"`
}
