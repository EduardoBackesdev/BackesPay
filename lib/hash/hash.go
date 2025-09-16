package lib

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func CheckPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// func main() {
// 	pass := "minhaSenhaSecreta"
// 	hash, _ := HashPassword(pass)
// 	fmt.Println("Hash:", hash)

// 	fmt.Println("Senha correta?", CheckPassword(hash, pass))
// 	fmt.Println("Senha incorreta?", CheckPassword(hash, "outraSenha"))
// }
