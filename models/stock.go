package models

import "time"

//Stock model is the model that is returned from the robin hood api
type Stock struct {
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
