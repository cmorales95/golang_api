package handlers

import (
	"encoding/json"
	"github.com/cmorales95/golang_api/crud/authorization"
	"github.com/cmorales95/golang_api/crud/models"
	"net/http"
)

type login struct {
	storage Storage
}

func newLogin(s Storage) *login {
	return &login{s}
}

func (l *login) login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response := newResponse(Error, "Method not allowed", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	data := models.Login{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		resp := newResponse(Error, "login struct is not valid", nil)
		responseJSON(w, http.StatusBadRequest, resp)
		return
	}

	// validate against bd
	if !isLoginValid(&data) {
		resp := newResponse(Error, "user or password is not valid", nil)
		responseJSON(w, http.StatusBadRequest, resp)
		return
	}

	// generate token for the session
	token, err := authorization.GenerateToken(&data)
	if err != nil {
		resp := newResponse(Error, "error, token was not generated", nil)
		responseJSON(w, http.StatusInternalServerError, resp)
		return
	}

	dataToken := map[string]string{"token":token}
	resp := newResponse(Message, "ok", dataToken)
	responseJSON(w, http.StatusOK, resp)
}

func isLoginValid(data *models.Login) bool {
	// dummy func
	return data.Email == "contact@mycompany" && data.Password == "123456"
}