package repositories

import "main/db"

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponseSucces struct {
	Message string
}

func LoginAccount(data LoginRequest) ([]LoginRequest, error) {

	var a []LoginRequest

	db, err := db.Conn()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT email, password from accounts where email = ?", data.Email)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var b LoginRequest
		if err := rows.Scan(&b.Email, &b.Password); err != nil {
			return nil, err
		}
		a = append(a, b)
	}

	defer db.Close()

	return a, nil

}
