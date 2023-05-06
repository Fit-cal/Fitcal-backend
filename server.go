package main

import (
	"fitcal-backend/infrastructure"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	infrastructure.Router(e)
	e.Logger.Fatal(e.Start(":8080"))
}
