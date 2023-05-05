package main

import (
	"fitcal-backend/routers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	routers.Router(e)
	e.Logger.Fatal(e.Start(":8080"))
}
