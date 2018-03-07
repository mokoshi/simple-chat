package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/signup", func(c echo.Context) error {
		return c.String(http.StatusOK, "Signup page")
	})

	e.GET("/login", func(c echo.Context) error {
		return c.String(http.StatusOK, "Login page")
	})

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Chat page")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
