package listeners

//Listener exposes methods to listen to stock symbols
type Listener interface {
	ListenToSymbol(symbol string)
}
