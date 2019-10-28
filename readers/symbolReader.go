package readers

//SymbolReaderType defines types of symbol readers
type SymbolReaderType int

const (
	//JSON file reader
	JSON = 0
	//InMemory reader
	InMemory = 1
)

//SymbolReader exposes methods to read stock symbols
type SymbolReader interface {
	GetSymbols() ([]string, error)
}

//NewReader creates a new symbol reader
func NewReader(readerType SymbolReaderType) (SymbolReader, error) {
	var reader SymbolReader
	var err error

	switch readerType {
	case JSON:
		reader = new(JSONReader)
	case InMemory:
		reader = new(InMemoryReader)
	}

	return reader, err
}
