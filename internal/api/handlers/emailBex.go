package handlers

import (
	"main/internal/api/services"
	"main/internal/repositories"

	"github.com/gin-gonic/gin"
)

func EmailBex(r *gin.Context) {

	var a repositories.EmailBexRequest

	if err := r.ShouldBindJSON(&a); err != nil {
		r.JSON(400, repositories.AccountResponseError{Error: err.Error()})
		return
	}

	result, err := services.EmailBex(a)
	if err != nil {
		r.JSON(400, repositories.EmailBexResponseError{Message: err.Error()})
		return
	}

	r.JSON(200, result)
	return

}
