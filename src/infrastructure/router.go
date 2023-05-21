package infrastructure

import (
	"fitcal-backend/controllers"
	"fitcal-backend/domain/interactor"
	"fitcal-backend/domain/repository"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Router(e *echo.Echo) {
	db, err := Connection()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	userRepo := repository.NewUserRepository(db)
	userInteractor := interactor.NewUserInteractor(userRepo)
	userController := controllers.NewUserController(userInteractor)

	api := e.Group("/api")
	api.GET("/users", func(c echo.Context) error {
		log.Print("GetUsers -->>")
		return c.JSON(http.StatusOK, userController.GetUser(c))
	})
	api.GET("/search/users", func(c echo.Context) error {
		log.Print("SearchUsers -->>")
		return c.JSON(http.StatusOK, userController.SearchUser(c))
	})
}
