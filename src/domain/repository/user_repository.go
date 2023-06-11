package repository

import (
	"fitcal-backend/domain/entities"
	"fitcal-backend/util"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// GetUsers gets all user
func (repository *UserRepository) GetUsers() (entities.Users, error) {
	db := repository.db
	var users []entities.User
	if err := db.Model(&users).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// SearchUsers searches for specific user
func (repository *UserRepository) SearchUsers(keyword string) ([]entities.User, error) {
	db := repository.db

	var users []entities.User
	likeString := util.LikeStringBuilder(keyword)
	// is the code below ok like this
	// can we not do anything better???
	if err := db.Model(&users).Where("first_name LIKE ?", likeString).Or("last_name LIKE ?", likeString).Or("email LIKE ?", likeString).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// CreateUser creates a new user if the user doesnot exist
func (repository *UserRepository) CreateUser(query entities.User) error {
	db := repository.db

	if err := db.Create(&query).Error; err != nil {
		return err
	}

	return nil
}
