package controllers

import (
	"fitcal-backend/domain"
	"fitcal-backend/interface/interactor"
	"log"
)

type TestController struct {
	testInteractor interactor.TestInteractor
}

func NewTestController(i interactor.TestInteractor) *TestController {
	return &TestController{i}
}

func (c *TestController) TestControllerTest() *domain.User {
	result, err := c.testInteractor.TestRepository.GetUserName()
	if err != nil {
		log.Fatal(err)
	}
	return result
}
