package listeners

//Listener exposes methods to listen to stock symbols
type Listener interface {
	ListenToSymbol(symbol string)
}

//NewListener creates a new listener
func NewListener() (Listener, error) {
	var lstnr Listener

	lstnr = new(RobinhoodListener)

	return lstnr, nil
}
