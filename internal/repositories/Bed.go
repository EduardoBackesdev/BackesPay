package repositories

import (
	"fmt"
	"main/db"

	"github.com/shopspring/decimal"
)

type BedRequest struct {
	Id           int             `json:"id"`
	Email_client string          `json:"email_client"`
	Balance      decimal.Decimal `json:"balance"`
}

func Bed(data BedRequest) error {

	db, err := db.Conn()
	if err != nil {
		return fmt.Errorf("Error with send bed!")
	}

	_, err_exec_2 := db.Exec(`UPDATE account_balance ab
	INNER JOIN accounts acc ON ab.account_id = acc.id
	SET ab.balance = ab.balance - ?
	WHERE acc.id = ? `, data.Balance, data.Id)

	if err_exec_2 != nil {
		return fmt.Errorf("Error with send bed!")
	}

	_, err_exec := db.Exec(`UPDATE account_balance ab
	INNER JOIN accounts acc ON ab.account_id = acc.id
	SET ab.balance = ab.balance + ?
	WHERE acc.email = ? `, data.Balance, data.Email_client)

	if err_exec != nil {
		return fmt.Errorf("Error with send bed!")
	}

	return nil

}
