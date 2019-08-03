package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	//Todo: pull top 100 movers and track these symbols concurrently
	var wg sync.WaitGroup
	wg.Add(2)

	go trackSymbol("MSFT")
	go trackSymbol("AAPL")

	wg.Wait()
}

//Tracking price for a given symbol
func trackSymbol(symbol string) {
	for {
		quotesEndpoint := "https://api.robinhood.com/marketdata/quotes/?symbols=" + symbol

		//Todo: ping oauth endpoint for token, find way to get
		var bearer = "Bearer " + "{ENTER YOUR TOKEN HERE}"

		req, error := http.NewRequest("GET", quotesEndpoint, nil)

		if error == nil {
			req.Header.Add("Authorization", bearer)
			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				log.Println("Error on response.\n[ERRO] -", err)
			}
			body, _ := ioutil.ReadAll(resp.Body)
			results := ResultsWrapper{}
			json.Unmarshal([]byte(body), &results)
			result := results.Results[0]
			fmt.Println(symbol, "Ask Price is", result.AskPrice)
		}
	}
}

//ResultsWrapper model
type ResultsWrapper struct {
	Results []struct {
		AskPrice                    string    `json:"ask_price"`
		AskSize                     int       `json:"ask_size"`
		BidPrice                    string    `json:"bid_price"`
		BidSize                     int       `json:"bid_size"`
		LastTradePrice              string    `json:"last_trade_price"`
		LastExtendedHoursTradePrice string    `json:"last_extended_hours_trade_price"`
		PreviousClose               string    `json:"previous_close"`
		AdjustedPreviousClose       string    `json:"adjusted_previous_close"`
		PreviousCloseDate           string    `json:"previous_close_date"`
		Symbol                      string    `json:"symbol"`
		TradingHalted               bool      `json:"trading_halted"`
		HasTraded                   bool      `json:"has_traded"`
		LastTradePriceSource        string    `json:"last_trade_price_source"`
		UpdatedAt                   time.Time `json:"updated_at"`
		Instrument                  string    `json:"instrument"`
	} `json:"results"`
}
