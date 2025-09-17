package handlers

import (
	"main/internal/api/services"
	"main/internal/repositories"

	"github.com/gin-gonic/gin"
)

func EmailBix(r *gin.Context) {

	var a repositories.EmailBixRequest

	if err := r.ShouldBindJSON(&a); err != nil {
		r.JSON(400, repositories.AccountResponseError{Error: err.Error()})
		return
	}

	result, err := services.EmailBix(a)
	if err != nil {
		r.JSON(400, repositories.EmailBixResponseError{Message: err.Error()})
		return
	}

	r.JSON(200, result)
	return

}
