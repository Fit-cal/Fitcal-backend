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

func (repository *UserRepository) GetUsers() ([]entities.User, error) {
	db := repository.db
	var users []entities.User
	if err := db.Model(&users).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repository *UserRepository) SearchUsers(keyword string) ([]entities.User, error) {
	db := repository.db

	var users []entities.User
	likeString := util.LikeStringBuilder(keyword)
	if err := db.Model(&users).Where("first_name LIKE ?", likeString).Or("last_name LIKE ?", likeString).Or("email LIKE ?", likeString).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
