package lib

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("Fail Hash")
	}
	return string(hash), err
}

func CheckPassword(hash, password string) (bool, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return false, errors.New("Invalid Password")
	}
	return true, nil
}

// func main() {
// 	pass := "minhaSenhaSecreta"
// 	hash, _ := HashPassword(pass)
// 	fmt.Println("Hash:", hash)

// 	fmt.Println("Senha correta?", CheckPassword(hash, pass))
// 	fmt.Println("Senha incorreta?", CheckPassword(hash, "outraSenha"))
// }
