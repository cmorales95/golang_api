package handlers

import (
	m "github.com/cmorales95/golang_api/crud/middlewares"
	"net/http"
)

func RoutePerson(mux *http.ServeMux, store Storage) {
	h := newPerson(store)
	mux.HandleFunc("/v1/persons/create", m.Log(m.Auth(h.create)))
	mux.HandleFunc("/v1/persons/get-all", m.Log(h.getAll))
	mux.HandleFunc("/v1/persons/update", m.Log(h.Update))
	mux.HandleFunc("/v1/persons/delete",m.Log(h.delete))
}

func RouteLogin(mux *http.ServeMux, storage Storage) {
	h := newLogin(storage)

	mux.HandleFunc("/v1/login", h.login)
}
