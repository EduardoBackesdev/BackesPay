package handlers

import (
	"fmt"
	"main/internal/api/services"
	"main/internal/repositories"

	"github.com/gin-gonic/gin"
)

func EmailBix(r *gin.Context) {

	var a repositories.EmailBixRequest

	if err := r.ShouldBindJSON(&a); err != nil {
		r.JSON(400, ErrorResponse{Message: fmt.Sprintf("Error with login account: %v", err)})
		return
	}

	result, err := services.EmailBix(a)
	if err != nil {
		r.JSON(400, ErrorResponse{Message: fmt.Sprintf("Error with login account: %v", err)})
		return
	}

	r.JSON(200, result)
	return

}
