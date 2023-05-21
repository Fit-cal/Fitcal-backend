package interactor

import (
	"fitcal-backend/domain/entities"
	"fitcal-backend/domain/repository/inputport"
	"strings"
)

type UserInteractor struct {
	userRepository inputport.UserRepositoryInputPort
}

func NewUserInteractor(userRepository inputport.UserRepositoryInputPort) *UserInteractor {
	return &UserInteractor{
		userRepository: userRepository,
	}
}

func (interactor *UserInteractor) GetUsers() ([]entities.User, error) {
	user, err := interactor.userRepository.GetUsers()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (interactor *UserInteractor) SearchUsers(keyword string) ([]entities.User, error) {
	keyword = strings.Trim(keyword, `"`)
	user, err := interactor.userRepository.SearchUsers(keyword)
	if err != nil {
		return nil, err
	}
	return user, nil
}
