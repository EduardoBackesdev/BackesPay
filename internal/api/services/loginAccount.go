package services

import (
	"fmt"
	"main/internal/repositories"
	hash "main/lib/hash"
	lib "main/lib/jwt"
)

func LoginAccount(data repositories.LoginRequest) (repositories.LoginResponseSucces, error) {

	result, err := repositories.LoginAccount(data)
	if err != nil {
		return repositories.LoginResponseSucces{}, fmt.Errorf("Error with Login account: %v", err)
	}

	_, err = hash.CheckPassword(result.Password, data.Password)
	if err != nil {
		return repositories.LoginResponseSucces{}, fmt.Errorf("Error check password: %v", err)
	}

	token, errToken := lib.CreateToken(result.Email)
	if errToken != nil {
		return repositories.LoginResponseSucces{}, fmt.Errorf("Error with create token: %v", errToken)
	}

	return repositories.LoginResponseSucces{
		Id:      result.Id,
		Message: "Login efetuado com sucesso",
		Email:   result.Email,
		Token:   token,
	}, nil

}
