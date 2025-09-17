package repositories

import (
	"database/sql"
	"main/db"
)

type query struct {
	Id       int
	Email    string
	Password string
}
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponseSucces struct {
	Id      int
	Email   string
	Token   string
	Message string
}

func LoginAccount(data LoginRequest) (query, error) {

	var a query

	db, err := db.Conn()
	if err != nil {
		return query{}, err
	}

	row := db.QueryRow("SELECT id, email, password from accounts where email = ?", data.Email)

	if err_row := row.Scan(&a.Id, &a.Email, &a.Password); err_row != nil {
		if err_row == sql.ErrNoRows {
			return query{}, err_row
		}
		return query{}, err_row
	}

	defer db.Close()

	return a, nil

}
