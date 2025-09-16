package handlers

import (
	"main/internal/api/services"
	"main/internal/repositories"

	"github.com/gin-gonic/gin"
)

type accountResponseError struct {
	Error error
}

type accountResponseSuccess struct {
	Message string
}

func CreateAccount(r *gin.Context) {

	var data repositories.AccountRequest

	if err := r.ShouldBindJSON(&data); err != nil {
		r.JSON(400, accountResponseError{Error: err})
		return
	}

	result, err := services.CreateAccount(data)
	if err != nil {
		r.JSON(400, accountResponseError{Error: err})
		return
	}

	r.JSON(200, accountResponseSuccess{Message: result.Message})
	return

}
