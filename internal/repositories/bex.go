package repositories

import "github.com/shopspring/decimal"

type bexRequest struct {
	Id           int
	Email_client string
	balance      decimal.Decimal
}

func Bex() {

}
