package repositories

import (
	"database/sql"
	"fmt"
	"main/db"
)

type EmailBixStruct struct {
	Name  string
	Email string
}

type EmailBixRequest struct {
	Email string
}

type EmailBixResponseSuccess struct {
	Name    string
	Email   string
	Message string
}

func EmailBix(data EmailBixRequest) (EmailBixStruct, error) {

	var a EmailBixStruct

	db, err := db.Conn()
	if err != nil {
		return EmailBixStruct{}, fmt.Errorf("Error with connection: %w", err)
	}

	row := db.QueryRow("SELECT email, name FROM accounts where email = ?", data.Email)
	if errRow := row.Scan(&a.Email, &a.Name); errRow != nil {
		if errRow == sql.ErrNoRows {
			return EmailBixStruct{}, fmt.Errorf("Error no rows to scan: %w", errRow)
		}
		return EmailBixStruct{}, fmt.Errorf("Error with scan row: %w", errRow)
	}

	return a, nil

}
