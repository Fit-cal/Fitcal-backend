package infrastructure

import (
	"fitcal-backend/interface/controllers"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Router(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		log.Print("health check")
		return c.JSON(http.StatusOK, "Health check")
	})
	api := e.Group("/api")

	api.GET("/", func(c echo.Context) error {
		log.Print("connection made")
		return c.JSON(http.StatusOK, controllers.TestControllerTest())
	})
}
