package services

import (
	"main/internal/repositories"
	lib "main/lib/hash"
)

func LoginAccount(data repositories.LoginRequest) (repositories.AccountResponseSuccess, error) {

	result, err := repositories.LoginAccount(data)
	if err != nil {
		return repositories.AccountResponseSuccess{}, err
	}

	_, err = lib.CheckPassword(result[0].Password, data.Password)
	if err != nil {
		return repositories.AccountResponseSuccess{}, err
	}

}
