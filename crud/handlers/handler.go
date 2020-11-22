package handlers

import "github.com/cmorales95/golang_api/crud/models"

//Storage .
type Storage interface {
	Create(person *models.Person) error
	Update(ID int, person *models.Person) error
	Delete(ID int) error
	GetByID(ID int) (models.Person, error)
	GetAll() (models.Persons, error)
}
