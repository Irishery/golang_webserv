package apiserver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type RawCurrency struct {
	Symbol         string  `json:"symbol"`
	Price24H       float64 `json:"price_24h"`
	Volume24H      float64 `json:"volume_24h"`
	LastTradePrice float64 `json:"last_trade_price"`
}

type CurrencyInfo struct {
	Price     float64 `json:"price"`
	Volume    float64 `json:"volume"`
	LastTrade float64 `json:"last_trade"`
}

func MakeRequest() map[string]*CurrencyInfo {
	var input_data = make([]RawCurrency, 0)
	var output_data = map[string]*CurrencyInfo{}

	resp, err := http.Get("https://api.blockchain.com/v3/exchange/tickers")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal(body, &input_data)

	if err != nil {
		fmt.Printf("err = %v\n", err)
	}

	for _, s := range input_data {
		output_data[s.Symbol] = &CurrencyInfo{
			Price:     s.Price24H,
			Volume:    s.Volume24H,
			LastTrade: s.LastTradePrice,
		}
	}

	return output_data

}
