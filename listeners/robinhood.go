package listeners

//RobinhoodListener listens to stock info from the robinhood api
type RobinhoodListener struct {
	StockInfoEndpointURL string
}

//ListenToSymbol listen and analyze stock info from rh api
func (l *RobinhoodListener) ListenToSymbol(symbol string) {

}
