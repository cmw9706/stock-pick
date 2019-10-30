package storage

import "stock-pick/models"

//Storage interface that is uses to store buying and selling data
type Storage interface {
	GetTransactions() ([]models.Transaction, error)
	GetTransaction(models.Transaction) (models.Transaction, error)
	SaveTransaction(models.Transaction) error
	SaveStock(models.Stock) error
	RemoveStock(models.Stock) error
	GetOwnedStock() ([]models.Stock, error)
}
