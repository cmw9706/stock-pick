package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

func main() {
	//Todo: pull top 100 movers and track these symbols concurrently
	var wg sync.WaitGroup
	wg.Add(2)

	changeChannel := make(chan ChangeMetadata)

	go listenForPriceChanges(changeChannel)
	go trackSymbol("AAPL", changeChannel)

	wg.Wait()
}

func convertFloatToFormattedString(input float64) string {
	return strconv.FormatFloat(input, 'f', 2, 32)
}

func listenForPriceChanges(change chan ChangeMetadata) {
	for {
		select {
		case delta := <-change:
			fmt.Println(delta.Symbol, " has changed from ", delta.OldPrice, " to ", delta.NewPrice, " => ", convertFloatToFormattedString(delta.Delta), "Original price: ", convertFloatToFormattedString(delta.OriginalPrice))
		}
	}
}

//Tracking price for a given symbol
func trackSymbol(symbol string, change chan ChangeMetadata) {
	quotesEndpoint := "https://api.robinhood.com/marketdata/quotes/?symbols=" + symbol
	var bearer = "Bearer " + ""
	oldPrice := 0.0
	orignalPrice := 0.0
	for {
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
			newPrice, parseErr := strconv.ParseFloat(result.AskPrice, 32)

			if parseErr != nil {
				fmt.Errorf("failed to parse ask price")
			}

			if newPrice != oldPrice {
				delta := newPrice - oldPrice
				changeInfo := ChangeMetadata{
					Symbol:        symbol,
					NewPrice:      convertFloatToFormattedString(newPrice),
					OldPrice:      convertFloatToFormattedString(oldPrice),
					Delta:         delta,
					OriginalPrice: orignalPrice,
				}

				change <- changeInfo
			}

			if oldPrice == 0.0 {
				orignalPrice = newPrice
			}

			oldPrice = newPrice

			//fmt.Println(symbol, "Ask Price is", result.AskPrice)
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

//ChangeMetadata model
type ChangeMetadata struct {
	Delta         float64
	OriginalPrice float64
	Symbol        string
	NewPrice      string
	OldPrice      string
}
