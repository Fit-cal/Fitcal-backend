package dataaccess

import (
	"fitcal-backend/config"
	"fitcal-backend/domain"
)

type TestRepository struct {
	TestRepository *config.ConfigDB
}

func (r *TestRepository) GetUserName() (*domain.User, error) {
	var result *domain.User
	result = &domain.User{
		Name:  "Sabir",
		Email: "sabirbarahi41@gmail.com",
	}
	return result, nil
}
