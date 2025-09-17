package handlers

import (
	"main/internal/api/services"
	"main/internal/repositories"

	"github.com/gin-gonic/gin"
)

func Bix(r *gin.Context) {

	var a repositories.BixRequest

	if err := r.ShouldBindJSON(&a); err != nil {
		r.JSON(400, repositories.BixResponseError{Message: err.Error()})
		return
	}

	result, err_service := services.Bix(a)
	if err_service != nil {
		r.JSON(400, repositories.BixResponseError{Message: err_service.Error()})
		return
	}

	r.JSON(200, result)
	return

}
