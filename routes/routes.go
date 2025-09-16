package routes

import (
	"main/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {

	r := gin.Default()
	r.GET("/create", handlers.CreateAccount)
	return r

}
