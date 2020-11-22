package handlers

import (
	"net/http"
)

func RoutePerson(mux *http.ServeMux, store Storage) {
	h := newPerson(store)
	mux.HandleFunc("/v1/persons/create", h.create)

}
