package main

import (
	"github.com/cmorales95/golang_api/crud/handlers"
	"github.com/cmorales95/golang_api/crud/storage"
	"log"
	"net/http"
)

func main() {
	storage := storage.NewMemory()
	mux := http.NewServeMux()

	handlers.RoutePerson(mux, storage)

	log.Println("application is running on port 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}

}
