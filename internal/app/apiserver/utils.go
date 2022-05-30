package apiserver

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Irishery/golang_webserv.git/internal/app/model"
)

func MakeRequest() []*model.Currency {
	var input_data = make([]*model.Currency, 0)

	resp, err := http.Get("https://api.blockchain.com/v3/exchange/tickers")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal(body, &input_data)

	if err != nil {
		log.Fatalln(err)
	}

	return input_data
}
