package controllers

import (
	"fitcal-backend/domain/entities"
	"fitcal-backend/domain/interactor/inputport"
	logger "fitcal-backend/log"
	"fitcal-backend/response"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

type UserController struct {
	userInteractor inputport.UserInteractorInputPort
}

func NewUserController(userInteractor inputport.UserInteractorInputPort) *UserController {
	return &UserController{
		userInteractor: userInteractor,
	}
}

// GetUser gets all the existing user
func (controller *UserController) GetUser(c echo.Context) []entities.User {
	result, err := controller.userInteractor.GetUsers()
	if err != nil {
		logger.Log(zerolog.ErrorLevel, err.Error())
		return nil
	}
	return result
}

// SearchUser searches the database for users
func (controller *UserController) SearchUser(c echo.Context) []entities.User {
	keyword := c.QueryParam("s")
	result, err := controller.userInteractor.SearchUsers(keyword)
	if err != nil {
		logger.Log(zerolog.PanicLevel, err.Error())
		return nil
	}
	return result
}

// CreateUser creates a new user if the user doesnot already exist
func (controller *UserController) CreateUser(c echo.Context) {
	var query entities.User
	var res response.CommonRes

	err := c.Bind(&query)
	if err != nil {
		res.Message = err.Error()
		logger.Log(zerolog.InfoLevel, err.Error())
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if err = controller.userInteractor.CreateUser(&query); err != nil {
		res.Message = err.Error()
		logger.Log(zerolog.ErrorLevel, res.Message)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res.Message = "user created!"
	logger.Log(zerolog.InfoLevel, res.Message)
	c.JSON(http.StatusOK, res)
	return
}
