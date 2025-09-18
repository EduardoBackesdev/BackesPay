package repositories

import (
	"fmt"
	"main/db"
	lib "main/lib/hash"
)

type AccountRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type AccountResponseError struct {
	Error string
}
type AccountResponseSuccess struct {
	Message string
}

// Fazer logica pra salvar imagem localmente referente ao usuario
// Fazer Password minimo 8 caracteres e required
// Fazer Email required
// Fazer Name required

func CreateAccount(data AccountRequest) (AccountResponseSuccess, error) {

	db, err := db.Conn()
	if err != nil {
		return AccountResponseSuccess{}, fmt.Errorf("Error with connection: %w", err)
	}
	defer db.Close()

	password, err := lib.HashPassword(data.Password)
	if err != nil {
		return AccountResponseSuccess{}, err
	}

	r, er1 := db.Exec("INSERT INTO accounts (status, email, password, name, id_group, image) VALUES (?,?,?,?,?,?)",
		1, data.Email, password, data.Name, 2, "no_image")

	if er1 != nil {
		return AccountResponseSuccess{}, fmt.Errorf("Error with exec query: %w", er1)
	}

	a, b := r.LastInsertId()
	if b != nil {
		return AccountResponseSuccess{}, fmt.Errorf("Error with get last id: %w", b)
	}

	_, er2 := db.Exec("INSERT INTO account_balance (account_id, balance) VALUES (?,?)", a, 0)

	if er2 != nil {
		return AccountResponseSuccess{}, fmt.Errorf("Error with exec query: %w", er2)
	}

	return AccountResponseSuccess{Message: "Conta criada com Sucesso"}, nil

}
