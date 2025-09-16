package services

import "main/internal/repositories"

func CreateAccount(data repositories.AccountRequest) (repositories.AccountResponseSuccess, error) {

	result, err := repositories.CreateAccount(data)
	if err != nil {
		return repositories.AccountResponseSuccess{}, err
	}

	return result, nil

}
