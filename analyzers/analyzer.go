package analyzers

import "stock-pick/models"

//Analyzer exposes functions to analyse stock day and make calls about wether to buy or sell
type Analyzer interface {
	ShouldBuy(models.Stock) (bool, error)
}
