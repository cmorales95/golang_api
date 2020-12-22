package handlers

import (
	"errors"
	"github.com/cmorales95/golang_api/crud/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type person struct {
	storage Storage
}

func newPerson(storage Storage) *person {
	return &person{storage: storage}
}

func (p *person) create(c echo.Context) error {
	data := models.Person{}
	err := c.Bind(&data)
	if err != nil {
		response := newResponse(Error, "error, json struct is wrong", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	err = p.storage.Create(&data)
	if err != nil {
		response := newResponse(Error, "error, a problem occurred saving a person", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := newResponse(Message, "person created successfully", nil)
	return c.JSON(http.StatusOK, response)
}

func (p *person) Update(c echo.Context) error {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := newResponse(Error, "error, 'ID' must be interger", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	data := models.Person{}
	err = c.Bind(&data)
	if err != nil {
		response := newResponse(Error, "error, json struct is worng", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	err = p.storage.Update(ID, &data)
	if err != nil {
		response := newResponse(Error, "error, a problem getting all persons", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := newResponse(Message, "person updated successfully", data)
	return c.JSON(http.StatusOK, response)
}

func (p *person) getAll(c echo.Context) error {
	data, err := p.storage.GetAll()
	if err != nil {
		response := newResponse(Message, "error, a problem all persons", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := newResponse(Message, "ok", data)
	return c.JSON(http.StatusOK, response)
}

func (p *person) delete(c echo.Context) error {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := newResponse(Error, "error, 'ID' must be interger", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	err = p.storage.Delete(ID)
	if errors.Is(err, models.ErrorIDPersonsDoesNotExist) {
		response := newResponse(Error, "error, 'ID' person does not exist", nil)
		return c.JSON(http.StatusNotFound, response)
	}

	if err != nil {
		response := newResponse(Error, "an error ocurr during delete a person ", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}
	response := newResponse(Message, "ok", nil)
	return c.JSON(http.StatusOK, response)
}

func (p *person) getByID(c echo.Context) error {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := newResponse(Error, "error, 'ID' must be interger", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	data, err := p.storage.GetByID(ID)
	if err != nil {
		response := newResponse(Error, "an error ocurr during delete a person ", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := newResponse(Message, "ok", data)
	return c.JSON(http.StatusOK, response)
}
