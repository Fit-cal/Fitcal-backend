package controllers

import (
	"fitcal-backend/domain"
)

func TestControllerTest() *domain.User {
	var result *domain.User
	result = &domain.User{
		Name:  "sabir",
		Email: "sabirbarahi41@gmail.com",
	}
	return result
}
