package services

import "main/internal/repositories"

func GetAccount(data repositories.GetAccountRequest) (repositories.GetAccountResponseSuccess, error) {

	result, err := repositories.GetAccount(data)
	if err != nil {
		return result, err
	}

	return result, nil

}
