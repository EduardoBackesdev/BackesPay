package repositories

import "github.com/shopspring/decimal"

type BedRequest struct {
	Id           int
	Email_client string
	Balance      decimal.Decimal
}

func Bed(data BedRequest) {

}
