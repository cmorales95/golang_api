package main

import (
	"github.com/cmorales95/golang_api/crud/authorization"
	"github.com/cmorales95/golang_api/crud/handlers"
	"github.com/cmorales95/golang_api/crud/storage"
	"log"
	"net/http"
)

func main() {
	// load certificates
	err := authorization.LoadCertificate("crud/cmd/certificates/app.rsa", "crud/cmd/certificates/app.rsa.pub")
	if err != nil {
		log.Fatalf("error, certificates were not loaded: %s", err)
	}

	// register routes
	storage := storage.NewMemory()
	mux := http.NewServeMux()
	handlers.RoutePerson(mux, storage)
	handlers.RouteLogin(mux, storage)

	log.Println("application is running on port 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}

}
