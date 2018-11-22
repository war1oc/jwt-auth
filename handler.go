package main

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

type handler struct{}

// Most of the code is taken from the echo guide
// https://echo.labstack.com/cookbook/jwt
func (h *handler) login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	// Check in your db if the user exists or not
	if username == "jon" && password == "password" {
		// Create token
		token := jwt.New(jwt.SigningMethodHS256)

		// Set claims
		// This is the information which frontend can use
		// The backend can also decode the token and get admin etc.
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = "Jon Doe"
		claims["admin"] = true
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		// Generate encoded token and send it as response.
		// The signing string should be secret (a generated UUID works too)
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]string{
			"token": t,
		})
	}

	return echo.ErrUnauthorized
}

// Most of the code is taken from the echo guide
// https://echo.labstack.com/cookbook/jwt
func (h *handler) private(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.String(http.StatusOK, "Welcome "+name+"!")
}
