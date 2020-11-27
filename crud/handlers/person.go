package handlers

import (
	"encoding/json"
	"errors"
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
		response := newResponse(Error, "error, a problem occurred saving a person", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	response := newResponse(Message, "person created successfully", nil)
	responseJSON(w, http.StatusCreated, response)
}

func (p *person) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		response := newResponse(Error, "error, method is not supported", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	ID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		response := newResponse(Error, "error, 'ID' must be interger", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	data := models.Person{}
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response := newResponse(Error, "error, json struct is worng", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	err = p.storage.Update(ID, &data)
	if err != nil {
		response := newResponse(Error, "error, a problem getting all persons", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}
	response := newResponse(Error, "error, person updated successfully", data)
	responseJSON(w, http.StatusInternalServerError, response)
}

func (p *person) getAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := newResponse(Error, "error, method is not supported", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	data, err := p.storage.GetAll()
	if err != nil {
		response := newResponse(Message, "error, a problem all persons", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Message, "ok", data)
	responseJSON(w, http.StatusOK, response)
}

func (p *person) delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		response := newResponse(Error, "error, method is not supported", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	ID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		response := newResponse(Error, "error, 'ID' must be interger", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	err = p.storage.Delete(ID)
	if errors.Is(err, models.ErrorIDPersonsDoesNotExist) {
		response := newResponse(Error, "error, 'ID' person does not exist", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	if err != nil {
		response := newResponse(Error, "an error ocurr during delete a person ", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}
	response := newResponse(Message, "ok", nil)
	responseJSON(w, http.StatusOK, response)
}
