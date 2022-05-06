package main

import (
	"echo-api-example/pkg/db"
	"echo-api-example/pkg/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	err := db.Connect()
	if err != nil {
		e.Logger.Error(err)
	}

	e.GET("/", hello)
	routes.InitUserRoutes(e)

	e.Logger.Fatal(e.Start(":8000"))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello!\n")
}
