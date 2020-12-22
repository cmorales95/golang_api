package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const url = "http://localhost:8080"

func main() {
	c := loginClient(url+"/v1/login", "contact@mycompany", "123456")
	fmt.Printf("%+v",c)
	person := Person{
		"Cristian",
		25,
		Communities{{"EDTeam"}, {"Golang en Español"}},
	}

	p := createPerson(url+"/v1/persons", c.Data.Token, &person)
	fmt.Printf("%v", p)
}

func httpClient(method, url, token string, body io.Reader) *http.Response {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Fatalf("Request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	client := http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Fatalf("request: %v", err)
	}
	return response
}

