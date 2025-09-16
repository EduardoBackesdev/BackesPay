package repositories

import (
	"main/db"
	lib "main/lib/hash"
)

type AccountRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Name     string `json:"name"`
}

type AccountRepositorieSuccess struct {
	Message string
}

// Fazer logica pra salvar imagem localmente
// Referente ao usuario

func CreateAccount(data AccountRequest) (AccountRepositorieSuccess, error) {

	db, err := db.Conn()
	if err != nil {
		return AccountRepositorieSuccess{}, err
	}

	password, _ := lib.HashPassword(data.Password)

	_, err = db.Exec("INSERT INTO accounts (status, login, password, email, name, id_group, image) VALUES (?,?,?,?,?,?,?)",
		1, data.Login, password, data.Email, data.Name, 2, "no_image")

	if err != nil {
		return AccountRepositorieSuccess{}, err
	}

	return AccountRepositorieSuccess{Message: "Conta criada com Sucesso"}, nil

}
