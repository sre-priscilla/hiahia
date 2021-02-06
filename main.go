package main

import (
	"net/http"

	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	e := echo.New()
	e.HideBanner = true
	e.Logger.SetLevel(log.INFO)

	p := prometheus.NewPrometheus("hiahia", nil)
	p.Use(e)

	e.GET("/routes", func(c echo.Context) error {
		return c.JSON(http.StatusOK, e.Routes())
	})
	e.GET("/say", func(c echo.Context) error {
		e.Logger.Error("someone say something")
		return c.JSON(http.StatusOK, echo.Map{"foo": "bar"})
	})
	e.GET("/smile", func(c echo.Context) error {
		e.Logger.Info("someone smiled")
		return c.JSON(http.StatusOK, echo.Map{"hello": "world"})
	})

	e.Logger.Fatal(e.Start(":8080"))
}
