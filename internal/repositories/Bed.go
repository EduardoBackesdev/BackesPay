package repositories

import "github.com/shopspring/decimal"

type BedRequest struct {
	Id           int             `json:"id"`
	Email_client string          `json:"email_client"`
	Balance      decimal.Decimal `json:"balance"`
}

func Bed(data BedRequest) {

}
