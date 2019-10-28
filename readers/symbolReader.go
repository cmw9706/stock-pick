package readers

//SymbolReader exposes methods to read stock symbols
type SymbolReader interface {
	GetSymbols() ([]string, error)
}
