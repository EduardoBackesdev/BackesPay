package services

import (
	"fmt"
	"main/internal/repositories"
)

func CreateAccount(data repositories.AccountRequest) (repositories.AccountResponseSuccess, error) {

	result, err := repositories.CreateAccount(data)
	if err != nil {
		return repositories.AccountResponseSuccess{}, fmt.Errorf("Error with create account: %v", err)
	}

	return result, nil

}
