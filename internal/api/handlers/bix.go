package handlers

import (
	"fmt"
	"main/internal/api/services"
	"main/internal/repositories"

	"github.com/gin-gonic/gin"
)

func Bix(r *gin.Context) {

	var a repositories.BixRequest

	if err := r.ShouldBindJSON(&a); err != nil {
		r.JSON(400, ErrorResponse{Message: fmt.Sprintf("Error with login account: %v", err)})
		return
	}

	result, err_service := services.Bix(a)
	if err_service != nil {
		r.JSON(400, ErrorResponse{Message: fmt.Sprintf("Error with login account: %v", err_service)})
		return
	}

	r.JSON(200, result)
	return

}
