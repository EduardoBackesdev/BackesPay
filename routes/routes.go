package routes

import (
	"main/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {

	r := gin.Default()
	r.POST("/create", handlers.CreateAccount)
	return r

}
