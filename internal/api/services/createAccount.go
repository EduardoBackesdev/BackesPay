package services

type AccountRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Name     string `json:"name"`
}

func CreateAccount(data AccountRequest) error {

}
