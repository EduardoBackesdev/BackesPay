package lib

import (
	"fmt"
	"main/internal/api/handlers"
	lib "main/lib/jwt"

	"github.com/gin-gonic/gin"
)

type unauthorized struct {
	Message string
}

func Auth() gin.HandlerFunc {

	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(401, handlers.ErrorResponse{Message: fmt.Sprintf("No Token")})
			c.Abort()
			return
		}
		tokenString = tokenString[len("Bearer "):]

		err := lib.VerifyToken(tokenString)
		if err != nil {
			c.JSON(401, handlers.ErrorResponse{Message: fmt.Sprintf("Invalid Token")})
			c.Abort()
			return
		}
		c.Next()
	}

}
