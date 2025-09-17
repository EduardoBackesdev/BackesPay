package repositories

import (
	"database/sql"
	"errors"
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

type EmailBixResponseError struct {
	Message string
}

func EmailBix(data EmailBixRequest) (EmailBixStruct, error) {

	var a EmailBixStruct

	db, err := db.Conn()
	if err != nil {
		return EmailBixStruct{}, err
	}

	row := db.QueryRow("SELECT email, name FROM accounts where email = ?", data.Email)
	if errRow := row.Scan(&a.Email, &a.Name); errRow != nil {
		if errRow == sql.ErrNoRows {
			return EmailBixStruct{}, errRow
		}
		return EmailBixStruct{}, errors.New("Erro ao ler query!")
	}

	return a, nil

}
