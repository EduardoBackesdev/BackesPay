package lib

import (
	"fmt"
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
			c.JSON(401, fmt.Errorf("No Token"))
			c.Abort()
			return
		}
		tokenString = tokenString[len("Bearer "):]

		err := lib.VerifyToken(tokenString)
		if err != nil {
			c.JSON(401, fmt.Errorf("Invalid Token"))
			c.Abort()
			return
		}
		c.Next()
	}

}
