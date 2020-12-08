package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"strconv"
)

func main() {
	e := echo.New()
	e.Use(middleware.Recover()) /*default of  echo*/
	e.GET("/", sayHello)
	e.GET("/dividir", divition)
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