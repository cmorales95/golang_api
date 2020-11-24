package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/cmorales95/golang_api/crud/models"
)

type person struct {
	storage Storage
}

func newPerson(storage Storage) *person {
	return &person{storage: storage}
}

func (p *person) create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response := newResponse(Error, "error, method is not supported", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	data := models.Person{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response := newResponse(Error, "error, json struct is wrong", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	err = p.storage.Create(&data)
	if err != nil {
		response := newResponse(Error, "error, a problem occurred saving a person",nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	response := newResponse(Error, "person created successfully", nil)
	responseJSON(w, http.StatusCreated, response)
}

func (p *person) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"ok": false, "message":"error, method is not supported"}`))
		return
	}

	ID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"ok": false, "message":"error, 'ID' must be integer"}`))

	}

	data := models.Person{}
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"ok": false, "message":"error, json struct is wrong"}`))
		return
	}

	err = p.storage.Update(ID, &data)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"ok": false, "message":"error, a problem getting all persons"}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{"ok": true, "message":"person updated successfully"}`))
}

func (p *person) getAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"ok": false, "message":"error, method is not supported"}`))
		return
	}

	resp, err := p.storage.GetAll()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"ok": false, "message":"error, a problem getting all persons"}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(&resp)
	w.WriteHeader(http.StatusOK)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"ok": false, "message":"error, a problem getting all persons"}`))
		return

	}
}
