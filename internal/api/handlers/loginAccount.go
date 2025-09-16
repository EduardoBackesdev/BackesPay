package handlers

import (
	"main/internal/api/services"
	"main/internal/repositories"

	"github.com/gin-gonic/gin"
)

type loginErrorResponse struct {
	Message string
}

func LoginAccount(c *gin.Context) {
	var data repositories.LoginRequest

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(400, loginErrorResponse{Message: err.Error()})
		return
	}

	result, err := services.LoginAccount(data)
	if err != nil {
		c.JSON(400, loginErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(200, result)
	return
}
