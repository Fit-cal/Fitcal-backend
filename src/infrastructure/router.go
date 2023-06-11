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
	// Gets all user information
	api.GET("/users", func(c echo.Context) error {
		log.Print("GetUsers -->>")
		return c.JSON(http.StatusOK, userController.GetUser(c))
	})

	// Gets user information according to the search query
	api.GET("/search/users", func(c echo.Context) error {
		log.Print("SearchUsers -->>")
		return c.JSON(http.StatusOK, userController.SearchUsers(c))
	})

	// Creates a new user if the user doesnot already exist
	api.POST("/create/user", func(c echo.Context) error {
		log.Print("CreateUser -->>")
		userController.CreateUser(c)
		return nil
	})
}
