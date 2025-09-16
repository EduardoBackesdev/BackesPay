package handlers

import (
	"main/internal/api/services"
	"main/internal/repositories"

	"github.com/gin-gonic/gin"
)

type accountResponseError struct {
	Error string
}

type accountResponseSuccess struct {
	Message string
}

func CreateAccount(r *gin.Context) {

	var data repositories.AccountRequest

	if err := r.ShouldBindJSON(&data); err != nil {
		r.JSON(400, accountResponseError{Error: err.Error()})
		return
	}

	if err := services.CreateAccount(data); err != nil {
		r.JSON(400, accountResponseError{Error: err.Error()})
		return
	}

	r.JSON(200, accountResponseSuccess{Message: "Conta criada com sucesso!"})
	return

}
