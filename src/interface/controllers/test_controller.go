package controllers

import (
	"fitcal-backend/domain"
	"fitcal-backend/interface/interactor"
)

type TestController struct {
	testInteractor interactor.TestInteractor
}

func NewTestController(i interactor.TestInteractor) *TestController {
	return &TestController{i}
}

func (c *TestController) GetUserName() domain.User {
	result := c.testInteractor.GetUserName()
	return result
}

func (c *TestController) GetAllUsers() []domain.User {
	result := c.testInteractor.GetAllUsers()
	return result
}
