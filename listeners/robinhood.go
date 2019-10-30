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
	var bearer = "Bearer " + "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJleHAiOjE1NzI0MjU4MzUsInRva2VuIjoieW44emJxQWdOMUM1UGJ1TlpmVG9NZVNvRlZCQXFIIiwidXNlcl9pZCI6ImFkNTk3OTFiLWZkMTAtNDJiOS1hYmI2LWQ4Mzc1MWEwZjI1YiIsImRldmljZV9oYXNoIjoiMDFlMGZhMGNjMDhmNDAxMTNkMmZlNzcxMTVmMGIyNzciLCJzY29wZSI6ImludGVybmFsIiwidXNlcl9vcmlnaW4iOiJVUyIsIm9wdGlvbnMiOmZhbHNlLCJsZXZlbDJfYWNjZXNzIjpmYWxzZX0.Bh7_cHEKi5LRDlXsziMA0SaBZ6hzqXbu3u3FU2ntvhVjlRomDnc438PQZDWwmV9G_QvS6d7B1YtlCyVNS938fUFcE_n4D6PnEBA573428dpAqjKu_gbWhxYsezwxr2cZFVJHvTrQI4xY4223ds2A3iOAeG7eLf1ajByELvX_CPIgnV_ecRSxYlAhBCl4aCebzeWvDTs97CwqpM0aQL5oTbsxzU2HWSi1kH9dxnj15uMP0BeTEtcqRpRSpuCGT92kA6I7zRfsIuiUZpiDhZLQZxeQI7dVc5DasIp-sWKd2p0vtf6nplWtw8AaAJKQ-2HyhHdYqYLProVpwsHTFKBJ3w"
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
