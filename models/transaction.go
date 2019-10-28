package models

//Transaction model that is uses to store data about purchases and sales of stock
type Transaction struct {
	Symbol string          `json:"symbol"`
	Type   TransactionType `json:"type"`
	Price  float32         `json:"price"`
}

// TransactionType defines available storage types
type TransactionType int

const (
	//Buy indicates a buy transaction
	Buy TransactionType = 0
	//Sell indicates a sell transaction
	Sell TransactionType = 0
)
