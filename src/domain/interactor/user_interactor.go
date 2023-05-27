package interactor

import (
	"errors"
	"fitcal-backend/domain/entities"
	"fitcal-backend/domain/repository/inputport"
	logger "fitcal-backend/log"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
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

func (interactor *UserInteractor) CreateUser(query *entities.User) error {
	validate := validator.New()
	user, err := interactor.userRepository.SearchUsers(query.Email)
	if err != nil {
		logger.Log(zerolog.ErrorLevel, err.Error())
		return err
	}

	uuid := uuid.New() //string uuid

	query = &entities.User{
		UserId:    uuid.String(),
		FirstName: query.FirstName,
		LastName:  query.LastName,
		Email:     query.Email,
	}

	err = validate.Struct(query)
	if err != nil {
		return err
	}

	count := len(user) //count the nuber of user with given email
	if count > 0 {
		err := errors.New("User Already Exists!")
		return err
	}

	err = interactor.userRepository.CreateUser(*query)

	if err != nil {
		return err
	}
	return nil
}
