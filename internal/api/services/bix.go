package services

import (
	"errors"
	"main/internal/repositories"

	"github.com/shopspring/decimal"
)

func Bix(data repositories.BixRequest) (repositories.BixResponseSucces, error) {

	x := repositories.GetAccountRequest{
		Id: data.Id,
	}

	result, err := repositories.GetAccount(x)
	if err != nil {
		return repositories.BixResponseSucces{}, err
	}

	if result.Balance.Sub(data.Balance).Cmp(decimal.NewFromInt(0)) < 0 {
		return repositories.BixResponseSucces{}, errors.New("Money insufficient to send Bex")
	}

	balanceSend := repositories.BixSendRequest{
		Id:           data.Id,
		Email_client: data.Email_client,
		Balance_send: result.Balance.Sub(data.Balance),
	}

	bix_result, err_bix := repositories.Bix(balanceSend)
	if err_bix != nil {
		return repositories.BixResponseSucces{}, err_bix
	}

	return bix_result, nil

}
