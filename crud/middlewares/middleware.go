package middlewares

import (
	"github.com/cmorales95/golang_api/crud/authorization"
	"github.com/labstack/echo/v4"
	"net/http"
)

type HandlerSign func(w http.ResponseWriter, r *http.Request)

// Auth validation
func Auth(f echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		// Validation Token
		_, err := authorization.ValidateToken(token)
		if err != nil {
			return  c.JSON(http.StatusBadRequest,  map[string]string{"error": "method not allowed"})
		}
		return f(c)
	}
}