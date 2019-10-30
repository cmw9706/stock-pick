package listeners

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"stock-pick/analyzers"
	"stock-pick/models"
	"strconv"
	"time"
)

//RobinhoodListener listens to stock info from the robinhood api
type RobinhoodListener struct {
	StockInfoEndpointURL string
}

//ListenToSymbol listen and analyze stock info from rh api
func (l *RobinhoodListener) ListenToSymbol(symbol string, activityChannel chan string) {
	fmt.Println("Listening to ", symbol, "...")

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
			results := models.Stock{}
			json.Unmarshal([]byte(body), &results)
			result := results.Results[0]
			newPrice, parseErr := strconv.ParseFloat(result.AskPrice, 32)

			if parseErr != nil {
				fmt.Println("failed to parse ask price")
				panic(parseErr)
			}

			if newPrice != oldPrice {
				delta := newPrice - oldPrice
				changeInfo := models.Delta{
					Symbol:        symbol,
					NewPrice:      convertFloatToFormattedString(newPrice),
					OldPrice:      convertFloatToFormattedString(oldPrice),
					Delta:         delta,
					OriginalPrice: orignalPrice,
				}

				fmt.Println(changeInfo)

				go analyzers.EvaluateDelta(results, changeInfo)
			}

			if oldPrice == 0.0 {
				orignalPrice = newPrice
			}

			oldPrice = newPrice

			fmt.Println(symbol, "Ask Price is", result.AskPrice)
			time.Sleep(3 * time.Second)
		}
	}
}

func convertFloatToFormattedString(input float64) string {
	return strconv.FormatFloat(input, 'f', 2, 32)
}
