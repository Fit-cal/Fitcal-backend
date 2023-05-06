package infrastructure

import (
	"fitcal-backend/config"
	"fitcal-backend/interface/controllers"
	"fitcal-backend/interface/interactor"
	"fitcal-backend/interface/repository"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

var (
	dbHandler config.ConfigDB
)

func getTestController() controllers.TestController {
	TestRepository := repository.NewTestRepo(dbHandler)
	testInteractor := interactor.NewTestInteractor(TestRepository)
	testController := controllers.NewTestController(testInteractor)
	return *testController
}

func Router(e *echo.Echo) {

	testRepository := getTestController()

	e.GET("/", func(c echo.Context) error {
		log.Print("health check")
		return c.JSON(http.StatusOK, "Health check")
	})
	api := e.Group("/api")

	api.GET("/", func(c echo.Context) error {
		log.Print("connection made")
		return c.JSON(http.StatusOK, testRepository.TestControllerTest())
	})
}
