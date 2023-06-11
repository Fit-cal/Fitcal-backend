package inputport

import (
	"fitcal-backend/domain/entities"
)

type UserInteractorInputPort interface {
	GetUsers() (entities.Users, error)
	SearchUsers(keyword string) (entities.Users, error)
	CreateUser(query *entities.User) error
}
