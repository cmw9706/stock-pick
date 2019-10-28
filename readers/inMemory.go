package readers

//InMemoryReader reads symbols in memory
type InMemoryReader struct {
}

//GetSymbols from memory
func (r *InMemoryReader) GetSymbols() ([]string, error) {
	return []string{"AAPL", "MSFT", "NSIT"}, nil
}
