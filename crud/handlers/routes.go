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
	//mux.HandleFunc("/v1/persons/get-all", m.Log(h.getAll))
	//mux.HandleFunc("/v1/persons/update", m.Log(h.Update))
	//mux.HandleFunc("/v1/persons/delete",m.Log(h.delete))
}

func RouteLogin(e *echo.Echo, storage Storage) {
	h := newLogin(storage)
	e.POST("/v1/login", h.login)
}
