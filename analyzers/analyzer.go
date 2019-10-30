package analyzers

import (
	"fmt"
	"math/rand"
	"stock-pick/brokers"
	"stock-pick/models"
	"stock-pick/storage"
)

//EvaluateDelta decides weather to by or sell stock
func EvaluateDelta(evalStock models.Stock, delta models.Delta) error {
	owned, err := stockIsOwned(evalStock)

	if err != nil {
		return err
	}

	if owned {
		evaluateToSell(evalStock)
	} else {
		evaluateToBuy(evalStock)
	}

	return nil
}

func stockIsOwned(stock models.Stock) (bool, error) {

	owned := false
	ownedStock, error := storage.DB.GetOwnedStock()
	if error != nil {
		return owned, error
	}

	for _, os := range ownedStock {
		if os.Results[0].Symbol == stock.Results[0].Symbol {
			owned = true
		}
	}

	return owned, nil
}

func evaluateToSell(symbolToSell models.Stock) {
	//Do some evaluation
	number := rand.Intn(100)

	if number > 83 {
		fmt.Println("Selling")

		broker, error := brokers.NewBroker()
		if error != nil {
			panic(error)
		}

		go broker.SellStock(symbolToSell)
	}
}

func evaluateToBuy(symbolToBuy models.Stock) {
	//Do some evaluation
	number := rand.Intn(100)

	if number < 50 {

		fmt.Println("BUY!")

		broker, error := brokers.NewBroker()

		if error != nil {
			panic(error)
		}

		go broker.BuyStock(symbolToBuy)
	}
}
