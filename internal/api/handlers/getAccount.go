package handlers

import (
	"main/internal/api/services"
	"main/internal/repositories"

	"github.com/gin-gonic/gin"
)

func GetAccount(r *gin.Context) {

	var a repositories.GetAccountRequest

	if err := r.ShouldBindJSON(&a); err != nil {
		r.JSON(400, repositories.GetAccountResponseError{Message: err.Error()})
		return
	}

	result, errGet := services.GetAccount(a)
	if errGet != nil {
		r.JSON(400, repositories.GetAccountResponseError{Message: errGet.Error()})
		return
	}

	r.JSON(200, result)
	return

}
