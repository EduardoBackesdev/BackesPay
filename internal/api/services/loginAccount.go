package services

import (
	"main/internal/repositories"
	hash "main/lib/hash"
	lib "main/lib/jwt"
)

func LoginAccount(data repositories.LoginRequest) (repositories.LoginResponseSucces, error) {

	result, err := repositories.LoginAccount(data)
	if err != nil {
		return repositories.LoginResponseSucces{}, err
	}

	_, err = hash.CheckPassword(result[0].Password, data.Password)
	if err != nil {
		return repositories.LoginResponseSucces{}, err
	}

	token, errToken := lib.CreateToken(result[0].Email)
	if errToken != nil {
		return repositories.LoginResponseSucces{}, errToken
	}

	return repositories.LoginResponseSucces{
		Message: "Login efetuado com sucesso",
		Email:   result[0].Email,
		Token:   token,
	}, nil

}
