package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)


func loginClient(url, email, password string) LoginResponse {
	login := Login{
		Email: email,
		Password: password,
	}

	data := bytes.NewBuffer([]byte{})
	err := json.NewEncoder(data).Encode(&login)
	if err != nil {
		log.Fatalf("Error en marshal: %v",err)
	}

	resp := httpClient(http.MethodPost, url,"", data)
	defer resp.Body.Close()


	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("error reading body request, %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("status code is wrong %q, %s", resp.StatusCode, body)
	}


	dataResponse := LoginResponse{}
	err = json.NewDecoder(bytes.NewReader(body)).Decode(&dataResponse)
	if err != nil {
		log.Fatalf("Error en el unmarshal del body")
	}

	return dataResponse
}
