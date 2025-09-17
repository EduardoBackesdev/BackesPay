package services

import (
	"main/internal/repositories"
)

func EmailBex(data repositories.EmailBexRequest) (repositories.EmailBexResponseSuccess, error) {

	result, err := repositories.EmailBex(data)
	if err != nil {
		return repositories.EmailBexResponseSuccess{}, err
	}

	a := repositories.EmailBexResponseSuccess{
		Name:    result.Name,
		Email:   result.Email,
		Message: "Email valid!",
	}

	return a, nil

}
