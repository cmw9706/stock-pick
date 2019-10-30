package storage

import (
	"stock-pick/models"
)

//DB mock database
var DB InMemoryStorage

//InMemoryStorage mock storage
type InMemoryStorage struct {
	Stocks       []models.Stock
	Transactions []models.Transaction
}

//GetTransactions mock implementation of getting transactions
func (s *InMemoryStorage) GetTransactions() ([]models.Transaction, error) {
	return DB.Transactions, nil
}

//GetTransaction mock implementation of getting single transaction
func (s *InMemoryStorage) GetTransaction(searchTransaction models.Transaction) (models.Transaction, error) {
	var transaction models.Transaction

	for _, t := range DB.Transactions {
		if t.ID == searchTransaction.ID {
			transaction = t
		}
	}

	return transaction, nil
}

//SaveTransaction saves a transaction to the mock db
func (s *InMemoryStorage) SaveTransaction(newTransaction models.Transaction) error {

	DB.Transactions = append(DB.Transactions, newTransaction)

	return nil
}

//GetOwnedStock returns owned stock from mock db
func (s *InMemoryStorage) GetOwnedStock() ([]models.Stock, error) {
	return DB.Stocks, nil
}
