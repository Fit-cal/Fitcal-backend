package entities

type User struct {
	UserId    string `json:"user_id" gorm:"user_id"`
	FirstName string `json:"first_name" gorm:"first_name" validate:"required"`
	LastName  string `json:"last_name" gorm:"last_name" validate:"required"`
	Email     string `json:"email" gorm:"email" validate:"required"`
}
