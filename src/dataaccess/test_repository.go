package dataaccess

import (
	"fitcal-backend/config"
	"fitcal-backend/domain"
	"fitcal-backend/interface/repository"
	"log"
)

type TestRepository struct {
	testRepository repository.TestRepository
}

func (r *TestRepository) GetUserName() (domain.User, error) {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalln(err)
	}
	user := domain.User{}

	if err := db.Model(&user).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *TestRepository) GetAllUsers() ([]domain.User, error) {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalln(err)
	}
	user := []domain.User{}
	if err := db.Model(&user).Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
