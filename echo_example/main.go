package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
	"strconv"
)

func main() {

	e := echo.New()
	e.Use(middleware.Recover()) /*default of  echo*/

	e.GET("/", sayHello)
	e.GET("/dividir", divition)

	//e.POST("/persons/create", create_person)
	//e.GET("/persons/get", get_person)
	//e.PUT("/persons/update", update_person)
	//e.DELETE("/persons/delete", delete_person)
	persons := e.Group("/persons") /*I have a group of persons*/
	persons.Use(middlewareLogPerson)
	persons.POST("", create_person)
	persons.GET("/:id", get_person)
	persons.PUT("/:id", update_person)
	persons.DELETE("/:id", delete_person)

	e.Logger.Fatal(e.Start(":8080"))
}

func sayHello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func divition(c echo.Context) error {
	d := c.QueryParam("d")
	f, _ := strconv.Atoi(d)
	r := 3000 / f
	return c.String(http.StatusOK, strconv.Itoa(r))
}

func create_person(c echo.Context) error {
	return c.String(http.StatusOK, "created ")
}

func get_person(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "got "+ id)
}

func update_person(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "updated " + id)
}

func delete_person(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "deleted " + id)
}

func middlewareLogPerson(f echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Println("petition realized to persons")
		return f(c)
	}
}