package handlers

import (
	"fmt"
	"main/internal/api/services"
	"main/internal/repositories"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Message string
}

func LoginAccount(c *gin.Context) {
	var data repositories.LoginRequest

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(400, gin.H{"Message": err})
		return
	}

	result, err := services.LoginAccount(data)
	if err != nil {
		c.JSON(400, ErrorResponse{Message: fmt.Sprintf("Error with login account: %v", err)})
		return
	}

	c.JSON(200, result)
	return
}
