package models

//Community model of struct for communities
type Community struct {
	Name string `json:"name"`
}

//Communities slice of community
type Communities []Community

//Person model of struct for persons
type Person struct {
	Name        string      `json:"name"`
	Age         uint        `json:"age"`
	Communities Communities `json:"communities"`
}

//Persons slice of person
type Persons []Person
