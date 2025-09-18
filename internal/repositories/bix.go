package repositories

import (
	"fmt"
	"main/db"

	"github.com/shopspring/decimal"
)

type BixRequest struct {
	Id           int
	Email_client string
	Balance      decimal.Decimal
}

type BixSendRequest struct {
	Id           int
	Email_client string
	Balance_send decimal.Decimal
}

type BixResponseSucces struct {
	Message string
}

type BixResponseError struct {
	Message string
}

func Bix(data BixSendRequest) (BixResponseSucces, error) {

	db, err := db.Conn()
	if err != nil {
		return BixResponseSucces{}, fmt.Errorf("Error with connection: %v", err)
	}

	_, err_exec := db.Exec(`UPDATE account_balance ab
	INNER JOIN accounts acc ON ab.account_id = acc.id
	SET ab.balance = ab.balance + ?
	WHERE acc.email = ? `, data.Balance_send, data.Email_client)

	if err_exec != nil {
		return BixResponseSucces{}, fmt.Errorf("Error with exec query: %v", err_exec)
	}

	return BixResponseSucces{"Bex send with success to account: " + data.Email_client}, nil

}
