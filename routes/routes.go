package routes

import (
	"main/internal/api/handlers"
	lib "main/lib/middlewares"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {

	r := gin.Default()

	//Websocket
	r.GET("ws", handlers.Ws)

	// Public
	r.POST("/create_account", handlers.CreateAccount)
	r.POST("/login", handlers.LoginAccount)

	// Auth
	auth := r.Group("/auth")
	auth.Use(lib.Auth())
	{
		auth.GET("/", handlers.GetAccount)

		bix := auth.Group("/account/debt")
		{
			bix.POST("/verify_email", handlers.EmailBix)
			bix.POST("/bix", handlers.Bix)
		}
	}

	return r

}
