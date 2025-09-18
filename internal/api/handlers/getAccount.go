package handlers

import (
	"fmt"
	"main/internal/api/services"
	"main/internal/repositories"

	"github.com/gin-gonic/gin"
)

func GetAccount(r *gin.Context) {

	var a repositories.GetAccountRequest

	if err := r.ShouldBindJSON(&a); err != nil {
		r.JSON(400, fmt.Errorf("Error with JSON request: %v", err))
		return
	}

	result, errGet := services.GetAccount(a)
	if errGet != nil {
		r.JSON(400, fmt.Errorf("Error with get account: %v", errGet))
		return
	}

	r.JSON(200, result)
	return

}
