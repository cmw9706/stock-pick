package brokers

import (
	"fmt"
	"math/rand"
	"stock-pick/models"
	"stock-pick/storage"
	"strconv"
	"time"
)

//RobinhoodBroker currently a move implementation of a stock buyer
type RobinhoodBroker struct {
}

//BuyStock mock implementation of buying stock
func (rb *RobinhoodBroker) BuyStock(stockToBuy models.Stock) error {
	s := stockToBuy.Results[0].Symbol
	p := stockToBuy.Results[0].BidPrice
	fmt.Println("Purchasing ", s, "...")

	time.Sleep(5 * time.Second)

	fmt.Println("Transaction complete!")

	go func() {
		id := strconv.Itoa(rand.Intn(1000000000))
		price, err := strconv.ParseFloat(p, 64)

		if err != nil {
			fmt.Println("failure to parse price")
		}

		newTransaction := models.Transaction{
			ID:     id,
			Symbol: s,
			Price:  price,
			Type:   models.Buy,
		}

		storage.DB.SaveTransaction(newTransaction)
	}()

	return nil
}

//SellStock mock implementation of selling stock
func (rb *RobinhoodBroker) SellStock(stockToSell models.Stock) error {
	s := stockToSell.Results[0].Symbol
	p := stockToSell.Results[0].BidPrice

	fmt.Println("Selling ", s, "...")

	time.Sleep(5 * time.Second)

	fmt.Println("Transaction complete!")

	go func() {
		id := strconv.Itoa(rand.Intn(1000000000))
		price, err := strconv.ParseFloat(p, 64)

		if err != nil {
			fmt.Println("failure to parse price")
		}

		newTransaction := models.Transaction{
			ID:     id,
			Symbol: s,
			Price:  price,
			Type:   models.Sell,
		}

		storage.DB.SaveTransaction(newTransaction)
	}()

	return nil
}
