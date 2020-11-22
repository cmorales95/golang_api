package models

import "errors"

var (
	//ErrPersonCanNotBeNil .
	ErrPersonCanNotBeNil = errors.New("person cannot be nul")
	//ErrorIDPersonsDoesNotExist .
	ErrorIDPersonsDoesNotExist = errors.New("person does not exists")
)
