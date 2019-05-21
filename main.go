package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	h := &handler{}

	e.POST("/login", h.login)

	e.GET("/private", h.private, isLoggedIn)

	e.GET("/admin", h.private, isLoggedIn, isAdmin)

	e.POST("/token", h.token)

	e.Logger.Fatal(e.Start(":1323"))
}
