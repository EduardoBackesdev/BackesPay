package handlers

import (
	"main/internal/api/services"
	"main/internal/repositories"

	"github.com/gin-gonic/gin"
)

func Bex(r *gin.Context) {

	var a repositories.BexRequest

	if err := r.ShouldBindJSON(&a); err != nil {
		r.JSON(400, repositories.BexResponseError{Message: err.Error()})
		return
	}

	result, err_service := services.Bex(a)
	if err_service != nil {
		r.JSON(400, repositories.BexResponseError{Message: err_service.Error()})
		return
	}

	r.JSON(200, result)
	return

}
