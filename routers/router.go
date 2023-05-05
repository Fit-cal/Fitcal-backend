package routers

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Name  string `json: "name" xml: "name"`
	Email string `json: "email" xml: "email"`
}

func Router(e *echo.Echo) {
	u := &User{
		Name:  "Sabir",
		Email: "sabirbarahi41@gmail.com",
	}
	e.GET("/", func(c echo.Context) error {
		log.Print("health check")
		return c.JSON(http.StatusOK, "Health check")
	})
	api := e.Group("/api")

	api.GET("/", func(c echo.Context) error {
		log.Print("connection made")
		return c.JSON(http.StatusOK, u)
	})
}
