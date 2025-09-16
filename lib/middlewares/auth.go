package lib

import (
	lib "main/lib/jwt"

	"github.com/gin-gonic/gin"
)

type unauthorized struct {
	Message string
}

func Auth() gin.HandlerFunc {

	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		if tokenString == "" {
			c.JSON(401, unauthorized{Message: "Token não enviado!"})
			c.Abort()
			return
		}
		tokenString = tokenString[len("Bearer "):]

		err := lib.VerifyToken(tokenString)
		if err != nil {
			c.JSON(401, unauthorized{Message: "Invalid token"})
			c.Abort()
			return
		}
		c.Next()
	}

}
