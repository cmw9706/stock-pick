package brokers

import (
	"stock-pick/models"
)

//Broker hooks into stock api to purchase stocks
type Broker interface {
	BuyStock(models.Stock) error
	SellStock(models.Stock) error
}

//NewBroker creates a new buyer to buy stock with
func NewBroker() (Broker, error) {
	var buyer Broker

	buyer = new(RobinhoodBroker)

	return buyer, nil
}
