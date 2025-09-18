package services

import (
	"fmt"
	"main/internal/repositories"

	"github.com/shopspring/decimal"
)

func Bix(data repositories.BixRequest) (repositories.BixResponseSucces, error) {

	x := repositories.GetAccountRequest{
		Id: data.Id,
	}

	result, err := repositories.GetAccount(x)
	if err != nil {
		return repositories.BixResponseSucces{}, fmt.Errorf("Error with get account: %w", err)
	}

	if result.Balance.Sub(data.Balance).Cmp(decimal.NewFromInt(0)) < 0 {
		return repositories.BixResponseSucces{}, fmt.Errorf("Error money insufficient")
	}

	balanceSend := repositories.BixSendRequest{
		Id:           data.Id,
		Email_client: data.Email_client,
		Balance_send: result.Balance.Sub(data.Balance),
	}

	bix_result, err_bix := repositories.Bix(balanceSend)
	if err_bix != nil {
		return repositories.BixResponseSucces{}, fmt.Errorf("Error with send bix: %w", err_bix)
	}

	return bix_result, nil

}
