package services

import (
	"fmt"
	"main/internal/repositories"
)

func EmailBix(data repositories.EmailBixRequest) (repositories.EmailBixResponseSuccess, error) {

	result, err := repositories.EmailBix(data)
	if err != nil {
		return repositories.EmailBixResponseSuccess{}, fmt.Errorf("Error with verify email bix: %w", err)
	}

	a := repositories.EmailBixResponseSuccess{
		Name:    result.Name,
		Email:   result.Email,
		Message: "Email valid!",
	}

	return a, nil

}
