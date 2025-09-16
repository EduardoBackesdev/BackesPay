package repositories

import (
	"database/sql"
	"main/db"
)

type AccountRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Name     string `json:"name"`
}

func CreateAccount(data AccountRequest) (*sql.DB, error) {

	db, err := db.Conn()
	if err != nil {
		return nil, err
	}

	result, err := db.Exec("INSERT INTO accounts (status, login, password, email, name, image) VALUES (?,?,?,?,?,?)")

}
