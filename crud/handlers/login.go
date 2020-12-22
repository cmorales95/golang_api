package handlers

import (
	"github.com/cmorales95/golang_api/crud/authorization"
	"github.com/cmorales95/golang_api/crud/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

type login struct {
	storage Storage
}

func newLogin(s Storage) *login {
	return &login{s}
}

func (l *login) login(c echo.Context) error {
	data := models.Login{}
	err := c.Bind(&data)
	if err != nil {
		resp := newResponse(Error, "login struct is not valid", nil)
		return c.JSON(http.StatusBadRequest, resp)
	}

	// validate against bd
	if !isLoginValid(&data) {
		resp := newResponse(Error, "user or password is not valid", nil)
		return c.JSON(http.StatusBadRequest, resp)
	}

	// generate token for the session
	token, err := authorization.GenerateToken(&data)
	if err != nil {
		resp := newResponse(Error, "error, token was not generated", nil)
		return c.JSON(http.StatusInternalServerError, resp)
	}

	dataToken := map[string]string{"token":token}
	resp := newResponse(Message, "ok", dataToken)
	return c.JSON(http.StatusOK, resp)
}

func isLoginValid(data *models.Login) bool {
	// dummy func
	return data.Email == "contact@mycompany" && data.Password == "123456"
}