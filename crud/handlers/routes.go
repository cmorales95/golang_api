package handlers

import (
	m "github.com/cmorales95/golang_api/crud/middlewares"
	"github.com/labstack/echo/v4"
)

func RoutePerson(e *echo.Echo, store Storage) {
	h := newPerson(store)
	person := e.Group("/v1/persons")
	person.Use(m.Auth)
	person.POST("", h.create)
	person.GET("", h.getAll)
	person.GET("/:id", h.getByID)
	person.PUT("/:id", h.Update)
	person.DELETE("/:id", h.delete)
}

func RouteLogin(e *echo.Echo, storage Storage) {
	h := newLogin(storage)
	e.POST("/v1/login", h.login)
}
