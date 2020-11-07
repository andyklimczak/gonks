package pkg

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Stock struct{
	Symbol string
	Price float64
	DollarChange float64
	PercentChange float64
}

type StockResponse struct {
	Close float64 `json:"c"`
	High float64 `json:"h"`
	Low float64 `json:"l"`
	Open float64 `json:"o"`
	PreviousClose float64 `json:"pc"`
}

func (m Stock) fetch(stockSymbols []string, apiKey string) (stocks []Stock) {
	for _, stockSymbol := range stockSymbols {
		url := fmt.Sprintf("https://finnhub.io/api/v1/quote?symbol=%s&token=%s", stockSymbol, apiKey)
		resp, err := http.Get(url)
		defer resp.Body.Close()
		if err != nil {
			panic(err)
		}
		var s = new(StockResponse)
		err = json.NewDecoder(resp.Body).Decode(&s)
		tmp := Stock{
			stockSymbol,
			s.Close,
			s.Close - s.PreviousClose,
			(s.Close - s.PreviousClose) / s.PreviousClose,
		}
		stocks = append(stocks, tmp)
	}
	return
}