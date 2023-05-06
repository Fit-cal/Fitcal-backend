package inputport

import "fitcal-backend/domain"

type TestRepository interface {
	GetUserName() (*domain.User, error)
}
