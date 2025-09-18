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
		return repositories.LoginResponseSucces{}, fmt.Errorf("Error with Login account: %w", err)
	}

	_, err = hash.CheckPassword(result.Password, data.Password)
	if err != nil {
		return repositories.LoginResponseSucces{}, fmt.Errorf("Error check password: %w", err)
	}

	token, errToken := lib.CreateToken(result.Email)
	if errToken != nil {
		return repositories.LoginResponseSucces{}, fmt.Errorf("Error with create token: %w", errToken)
	}

	return repositories.LoginResponseSucces{
		Id:      result.Id,
		Message: "Login efetuado com sucesso",
		Email:   result.Email,
		Token:   token,
	}, nil

}
