package controllers

import (
	"fitcal-backend/domain/entities"
	"fitcal-backend/domain/interactor/inputport"
	"log"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userInteractor inputport.UserInteractorInputPort
}

func NewUserController(userInteractor inputport.UserInteractorInputPort) *UserController {
	return &UserController{
		userInteractor: userInteractor,
	}
}

func (controller *UserController) GetUser(c echo.Context) []entities.User {
	result, err := controller.userInteractor.GetUsers()
	if err != nil {
		log.Print(err)
		return nil
	}
	return result
}

func (controller *UserController) SearchUser(c echo.Context) []entities.User {
	keyword := c.QueryParam("s")
	result, err := controller.userInteractor.SearchUsers(keyword)
	if err != nil {
		log.Panic(err)
		return nil
	}
	return result
}
