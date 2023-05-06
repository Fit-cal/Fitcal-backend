package repository

import (
	"fitcal-backend/domain"
	"fitcal-backend/interface/repository/inputport"
)

type TestRepository struct {
	testRepository inputport.TestRepository
}

func NewTestRepo(handler inputport.TestRepository) TestRepository {
	return TestRepository{handler}
}

func (r TestRepository) GetUserName() (*domain.User, error) {
	return r.testRepository.GetUserName()
}
