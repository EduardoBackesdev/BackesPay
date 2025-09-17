package repositories

import (
	"database/sql"
	"errors"
	"main/db"
)

type EmailBexStruct struct {
	Name  string
	Email string
}

type EmailBexRequest struct {
	Email string
}

type EmailBexResponseSuccess struct {
	Name    string
	Email   string
	Message string
}

type EmailBexResponseError struct {
	Message string
}

func EmailBex(data EmailBexRequest) (EmailBexStruct, error) {

	var a EmailBexStruct

	db, err := db.Conn()
	if err != nil {
		return EmailBexStruct{}, err
	}

	row := db.QueryRow("SELECT email, name FROM accounts where email = ?", data.Email)
	if errRow := row.Scan(&a.Email, &a.Name); errRow != nil {
		if errRow == sql.ErrNoRows {
			return EmailBexStruct{}, errRow
		}
		return EmailBexStruct{}, errors.New("Erro ao ler query!")
	}

	return a, nil

}
