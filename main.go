package main

import (
	"net/http"

	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	p := prometheus.NewPrometheus("hiahia", nil)
	p.Use(e)

	e.GET("/routes", func(c echo.Context) error {
		return c.JSON(http.StatusOK, e.Routes())
	})

	e.Logger.Fatal(e.Start(":8080"))
}
