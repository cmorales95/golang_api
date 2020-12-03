package handlers

import (
	"github.com/cmorales95/golang_api/crud/middlewares"
	"net/http"
)

func RoutePerson(mux *http.ServeMux, store Storage) {
	h := newPerson(store)
	mux.HandleFunc("/v1/persons/create", middlewares.Log(h.create))
	mux.HandleFunc("/v1/persons/get-all", middlewares.Log(h.getAll))
	mux.HandleFunc("/v1/persons/update", middlewares.Log(h.Update))
	mux.HandleFunc("/v1/persons/delete",middlewares.Log(h.delete))
}
