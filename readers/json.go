package readers

//JSONReader reads symbols from a json file
type JSONReader struct {
}

//GetSymbols from json file
func (r *JSONReader) GetSymbols() ([]string, error) {
	return []string{"test", "test", "test"}, nil
}
