package handlers

import (
	"fmt"
	"main/internal/api/services"
	"main/internal/repositories"

	"github.com/gin-gonic/gin"
)

func CreateAccount(r *gin.Context) {

	var data repositories.AccountRequest

	if err := r.ShouldBindJSON(&data); err != nil {
		r.JSON(400, ErrorResponse{Message: fmt.Sprintf("Error with login account: %v", err)})
		return
	}

	result, err := services.CreateAccount(data)
	if err != nil {
		r.JSON(400, ErrorResponse{Message: fmt.Sprintf("Error with login account: %v", err)})
		return
	}

	r.JSON(200, result)
	return

}
