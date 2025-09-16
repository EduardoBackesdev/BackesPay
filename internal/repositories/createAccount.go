package repositories

import (
	"main/db"
	lib "main/lib/hash"
)

type AccountRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type AccountResponseSuccess struct {
	Message string
}

// Fazer logica pra salvar imagem localmente
// Referente ao usuario

func CreateAccount(data AccountRequest) (AccountResponseSuccess, error) {

	db, err := db.Conn()
	if err != nil {
		return AccountResponseSuccess{}, err
	}

	password, err := lib.HashPassword(data.Password)
	if err != nil {
		return AccountResponseSuccess{}, err
	}

	_, err = db.Exec("INSERT INTO accounts (status, email, password, , name, id_group, image) VALUES (?,?,?,?,?,?)",
		1, data.Email, password, data.Name, 2, "no_image")

	if err != nil {
		return AccountResponseSuccess{}, err
	}

	defer db.Close()
	return AccountResponseSuccess{Message: "Conta criada com Sucesso"}, nil

}
