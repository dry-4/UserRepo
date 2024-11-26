package repositories

import (
	"userRepo/databses"
	"userRepo/models"
)

type UserRepository interface {
	FindAll() ([]models.User, error)
	FindById(id int) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (u *userRepository) FindAll() ([]models.User, error) {
	var users []models.User
	err := databses.GetDB().Find(&users).Error
	return users, err
}

func (u *userRepository) FindById(id int) (*models.User, error) {
	var user models.User
	err := databses.GetDB().First(&user, id).Error
	return &user, err
}

func (u *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := databses.GetDB().Where("email = ?", email).First(&user).Error
	return &user, err
}
