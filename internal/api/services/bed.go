package services

import (
	"fmt"
	"main/internal/repositories"
	lib "main/lib/goroutines"

	"github.com/shopspring/decimal"
)

func Bed(data repositories.BedRequest) error {

	a := repositories.EmailBixRequest{
		Email: data.Email_client,
	}

	b := repositories.GetAccountRequest{
		Id: data.Id,
	}

	result, err := EmailBix(a)
	if err != nil {
		return fmt.Errorf("Error with verify email bed: %w", err)

	}

	result_account, err_ := repositories.GetAccount(b)
	if err_ != nil {
		return fmt.Errorf("Error with get account: %w", err_)
	}

	if result_account.Balance.Sub(data.Balance).Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("Error money insufficient")
	}

	x := repositories.BedRequest{
		Id:           data.Id,
		Email_client: result.Email,
		Balance:      data.Balance,
	}

	go lib.Routine_bed(x)

	return nil

}
