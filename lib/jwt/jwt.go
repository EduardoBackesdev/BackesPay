package lib

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func getSecretKey() ([]byte, error) {
	if err := godotenv.Load("../../.env"); err != nil {
		return nil, errors.New("Fail get secret key")
	}
	var secret = []byte(os.Getenv("JWT"))
	return secret, nil
}

func CreateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})
	result, err := getSecretKey()
	if err != nil {
		return "", err
	}
	tokenString, err := token.SignedString(result)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		result, err := getSecretKey()
		if err != nil {
			return "", err
		}
		return result, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
