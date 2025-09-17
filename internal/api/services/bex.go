package services

import (
	"errors"
	"main/internal/repositories"

	"github.com/shopspring/decimal"
)

func Bex(data repositories.BexRequest) (repositories.BexResponseSucces, error) {

	x := repositories.GetAccountRequest{
		Id: data.Id,
	}

	result, err := repositories.GetAccount(x)
	if err != nil {
		return repositories.BexResponseSucces{}, err
	}

	if result.Balance.Sub(data.Balance).Cmp(decimal.NewFromInt(0)) < 0 {
		return repositories.BexResponseSucces{}, errors.New("Money insufficient to send Bex")
	}

	balanceSend := repositories.BexSendRequest{
		Id:           data.Id,
		Email_client: data.Email_client,
		Balance_send: result.Balance.Sub(data.Balance),
	}

	bex_result, err_bex := repositories.Bex(balanceSend)
	if err_bex != nil {
		return repositories.BexResponseSucces{}, err_bex
	}

	return bex_result, nil

}
