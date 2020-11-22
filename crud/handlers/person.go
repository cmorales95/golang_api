package handlers

import (
	"encoding/json"
	"github.com/cmorales95/golang_api/crud/models"
	"net/http"
)

type person struct {
	storage Storage
}

func newPerson(storage Storage) *person {
	return &person{storage: storage}
}

func (p *person) create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"ok": false, "message":"error, method is not supported"}`))
		return
	}

	data := models.Person{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"ok": false, "message":"error, json struct is wrong"}`))
		return
	}

	err = p.storage.Create(&data)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"ok": false, "message":"error, a problem occurred saving a person"}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte(`{"ok": true, "message":"person created successfully"}`))
}
