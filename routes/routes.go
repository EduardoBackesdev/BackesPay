package routes

import (
	"main/internal/api/handlers"
	lib "main/lib/middlewares"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {

	r := gin.Default()

	// Public
	r.POST("/create_account", handlers.CreateAccount)
	r.POST("/login", handlers.LoginAccount)

	// Auth
	auth := r.Group("/auth")
	auth.Use(lib.Auth())

	return r

}
