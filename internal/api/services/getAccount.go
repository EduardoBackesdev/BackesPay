package services

import (
	"fmt"
	"main/internal/repositories"
)

func GetAccount(data repositories.GetAccountRequest) (repositories.GetAccountResponseSuccess, error) {

	result, err := repositories.GetAccount(data)
	if err != nil {
		return result, fmt.Errorf("Error with get account: %w", err)
	}

	return result, nil

}
