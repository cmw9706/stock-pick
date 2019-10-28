package buyers

import (
	"stock-pick/models"
)

//Buyer hooks into stock api to purchase stocks
type Buyer interface {
	BuyStock(models.Stock) error
}
