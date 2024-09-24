package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/users", func(c echo.Context) error {
		user := new(User)

		if err := c.Bind(user); err != nil {
			return c.String(http.StatusBadRequest, "Invalid input")
		}

		return c.JSON(http.StatusOK, user)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
