package services

import (
	"main/internal/repositories"
)

func EmailBix(data repositories.EmailBixRequest) (repositories.EmailBixResponseSuccess, error) {

	result, err := repositories.EmailBix(data)
	if err != nil {
		return repositories.EmailBixResponseSuccess{}, err
	}

	a := repositories.EmailBixResponseSuccess{
		Name:    result.Name,
		Email:   result.Email,
		Message: "Email valid!",
	}

	return a, nil

}
