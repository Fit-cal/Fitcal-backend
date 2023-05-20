package repository

import (
	"fitcal-backend/domain"
	"fitcal-backend/interface/repository/inputport"
	"log"
)

type TestRepository struct {
	testRepository inputport.TestRepository
}

func NewTestRepo(handler inputport.TestRepository) *TestRepository {
	return &TestRepository{handler}
}

func (r *TestRepository) GetUserName() (domain.User, error) {
	user, err := r.testRepository.GetUserName()
	if err != nil {
		log.Fatalln(err)
	}
	return user, nil
}

func (r *TestRepository) GetAllUsers() ([]domain.User, error) {
	user, err := r.testRepository.GetAllUsers()
	if err != nil {
		log.Fatalln(err)
	}
	return user, nil
}
