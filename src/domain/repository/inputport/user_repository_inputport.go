package inputport

import "fitcal-backend/domain/entities"

type UserRepositoryInputPort interface {
	GetUsers() ([]entities.User, error)
	SearchUsers(keyword string) ([]entities.User, error)
}
