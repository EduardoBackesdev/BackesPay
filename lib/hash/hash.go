package lib

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("Fail hash: %v", err)
	}
	return string(hash), fmt.Errorf("Error with Hash: %v", err)
}

func CheckPassword(hash, password string) (bool, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return false, fmt.Errorf("Invalid Password!")
	}
	return true, nil
}
