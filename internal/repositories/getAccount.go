package repositories

import (
	"database/sql"
	"fmt"
	"main/db"

	"github.com/shopspring/decimal"
)

// Struct para Response de sucesso do Handler
type GetAccountResponseSuccess struct {
	Id      int
	Email   string
	Name    string
	Balance decimal.Decimal
}

// Struct para Response de erro do Handler
type GetAccountResponseError struct {
	Message string
}

// Struct para request do metodo
type GetAccountRequest struct {
	Id int
}

func GetAccount(data GetAccountRequest) (GetAccountResponseSuccess, error) {

	var a GetAccountResponseSuccess

	db, err := db.Conn()
	if err != nil {
		return GetAccountResponseSuccess{}, fmt.Errorf("Error with connection: %v", err)
	}

	row := db.QueryRow(`SELECT acc.id, acc.email, acc.name, ab.balance 
	from accounts acc
	inner join account_balance ab on (ab.account_id = acc.id)
	where acc.id = ?`, data.Id)
	if errRow := row.Scan(&a.Id, &a.Email, &a.Name, &a.Balance); errRow != nil {
		if errRow == sql.ErrNoRows {
			return GetAccountResponseSuccess{}, fmt.Errorf("Error no rows to scan: %v", errRow)
		}

		return GetAccountResponseSuccess{}, fmt.Errorf("Error with scan row: %v", errRow)
	}

	defer db.Close()

	return a, nil

}
