package inputport

import (
	"fitcal-backend/domain/entities"
)

type UserInteractorInputPort interface {
	GetUsers() ([]entities.User, error)
	SearchUsers(keyword string) ([]entities.User, error)
}
