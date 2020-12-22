package main

import (
	"github.com/cmorales95/golang_api/crud/authorization"
	"github.com/cmorales95/golang_api/crud/handlers"
	"github.com/cmorales95/golang_api/crud/storage"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	// load certificates
	err := authorization.LoadCertificate("crud/cmd/certificates/app.rsa", "crud/cmd/certificates/app.rsa.pub")
	if err != nil {
		log.Fatalf("error, certificates were not loaded: %s", err)
	}

	// register routes
	storage := storage.NewMemory()

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	handlers.RoutePerson(e, storage)
	handlers.RouteLogin(e, storage)
	log.Println("application is running on port 8080")
	err = e.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}

}
