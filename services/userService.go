package services

import (
	"userRepo/models"
	"userRepo/repositories"
)

type UserService interface {
	GetUsersService() ([]models.User, error)
	GetUserByIdService(id int) (*models.User, error)
	GetUserByEmailService(email string) (*models.User, error)
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repositories repositories.UserRepository) UserService {
	return &userService{repo: repositories}
}

func (s *userService) GetUsersService() ([]models.User, error) {
	return s.repo.FindAll()
}

func (s *userService) GetUserByIdService(id int) (*models.User, error) {
	return s.repo.FindById(id)
}

func (s *userService) GetUserByEmailService(email string) (*models.User, error) {
	return s.repo.FindByEmail(email)
}
