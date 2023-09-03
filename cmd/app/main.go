package main

import (
	"net/http"

	"github.com/Kitsuya0828/gqlgen-ent-clean-architecture-boilerplate/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config.ReadConfig(config.ReadConfigOption{})
	e := echo.New()

	e.Use(middleware.Recover())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome!")
	})

	e.Logger.Fatal(e.Start(":" + config.C.Server.Address))
}
