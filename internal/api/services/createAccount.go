package services

import "main/internal/repositories"

func CreateAccount(data repositories.AccountRequest) (repositories.AccountRepositorieSuccess, error) {

	result, err := repositories.CreateAccount(data)
	if err != nil {
		return repositories.AccountRepositorieSuccess{}, err
	}

	return result, nil

}
