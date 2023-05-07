package infrastructure

import (
	"fitcal-backend/dataaccess"
	"fitcal-backend/interface/controllers"
	"fitcal-backend/interface/interactor"
	"fitcal-backend/interface/repository"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

var (
	DB dataaccess.TestRepository
)

func getTestController() controllers.TestController {
	TestRepository := repository.NewTestRepo(&DB)
	testInteractor := interactor.NewTestInteractor(TestRepository)
	testController := controllers.NewTestController(testInteractor)
	return *testController
}

func Router(e *echo.Echo) {

	testController := getTestController()

	e.GET("/", func(c echo.Context) error {
		log.Print("health check")
		return c.JSON(http.StatusOK, "Health check")
	})
	api := e.Group("/api")

	api.GET("/", func(c echo.Context) error {
		log.Print("get user name -->>")
		return c.JSON(http.StatusOK, testController.GetUserName())
	})
	api.GET("/users", func(c echo.Context) error {
		log.Print("get all users -->>")
		return c.JSON(http.StatusOK, testController.GetAllUsers())
	})
}
