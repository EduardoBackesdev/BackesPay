package test

import (
	"bytes"
	"encoding/json"
	"main/internal/repositories"
	"main/routes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
)

// Test Integration
func TestCreateAccount(t *testing.T) {

	t.Run("testing success 200", func(tt *testing.T) {
		request := repositories.AccountRequest{
			Email:    "eduardoteste@hotmail.com",
			Password: "Eduardosenha123",
			Name:     "Eduardinho",
		}
		data, _ := json.Marshal(request)
		r := routes.Router()
		w := httptest.NewRecorder()
		req, err := http.NewRequest("POST", "/create_account", bytes.NewBuffer(data))
		if err != nil {
			t.Error("error %w", err)
		}
		r.ServeHTTP(w, req)
		t.Log(w.Body.String())
		assert.Equal(t, 200, w.Code)
	})

}
