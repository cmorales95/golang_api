package storage

import (
	"fmt"
	"github.com/cmorales95/golang_api/crud/models"
)

//Memory model struct
type Memory struct {
	currentID int
	Persons   map[int]models.Person
}

//NewMemory return an instance of memory
func NewMemory() *Memory {
	persons := make(map[int]models.Person)
	return &Memory{
		currentID: 0,
		Persons:   persons,
	}
}

func (m *Memory) Create(person *models.Person) error {
	if person == nil {
		return models.ErrPersonCanNotBeNil
	}

	m.currentID++
	m.Persons[m.currentID] = *person

	return nil
}

func (m *Memory) Update(ID int, person *models.Person) error {
	if person == nil {
		return models.ErrPersonCanNotBeNil
	}

	if _, ok := m.Persons[ID]; !ok {
		return fmt.Errorf("ID: %d: %w", ID, models.ErrorIDPersonsDoesNotExist)
	}

	m.Persons[ID] = *person
	return nil
}

func (m *Memory) Delete(ID int) error {
	if _, ok := m.Persons[ID]; !ok {
		return fmt.Errorf("ID: %d: %w", ID, models.ErrorIDPersonsDoesNotExist)
	}
	delete(m.Persons, ID)

	return nil
}

func (m *Memory) GetByID(ID int) (models.Person, error) {
	person, ok := m.Persons[ID]
	if !ok {
		return person, fmt.Errorf("ID: %d: %w", ID, models.ErrorIDPersonsDoesNotExist)
	}
	return person, nil
}

func (m *Memory) GetAll() (models.Persons, error) {
	var result models.Persons
	for _, v := range m.Persons {
		result = append(result, v)
	}
	return result, nil
}
