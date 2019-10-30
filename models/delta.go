package models

//Delta model for storing change data
type Delta struct {
	Delta         float64
	OriginalPrice float64
	Symbol        string
	NewPrice      string
	OldPrice      string
}
