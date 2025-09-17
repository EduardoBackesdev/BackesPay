package repositories

import (
	"main/db"

	"github.com/shopspring/decimal"
)

type BexRequest struct {
	Id           int
	Email_client string
	Balance      decimal.Decimal
}

type BexSendRequest struct {
	Id           int
	Email_client string
	Balance_send decimal.Decimal
}

type BexResponseSucces struct {
	Message string
}

type BexResponseError struct {
	Message string
}

func Bex(data BexSendRequest) (BexResponseSucces, error) {

	db, err := db.Conn()
	if err != nil {
		return BexResponseSucces{}, err
	}

	_, err_exec := db.Exec(`UPDATE account_balance ab
	INNER JOIN accounts acc ON ab.account_id = acc.id
	SET ab.balance = ab.balance + ?
	WHERE acc.email = ? `, data.Balance_send, data.Email_client)

	if err_exec != nil {
		return BexResponseSucces{}, err_exec
	}

	return BexResponseSucces{"Bex send with success to account: " + data.Email_client}, nil

}
