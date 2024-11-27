package services

import (
	"errors"
	"testing"
	"userRepo/models"
	"userRepo/repositories"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	mockRepo = new(repositories.MockUserRepository)
	service  = NewUserService(mockRepo)
)

func TestGetUsersServiceSuccess(t *testing.T) {
	t.Run("GetUsersService Success", func(t *testing.T) {
		expectedUsers := []models.User{
			{
				Id:    1,
				Name:  "Muskan",
				Email: "muskan@gmail.com",
				Age:   25,
			},
		}

		mockRepo.On("FindAll").Return(expectedUsers, nil)

		users, err := service.GetUsersService()

		assert.NoError(t, err)
		assert.Equal(t, expectedUsers, users)
		mockRepo.AssertExpectations(t)
	})
}

func TestGetUsersServiceError(t *testing.T) {
	mockRepo = new(repositories.MockUserRepository)
	service = NewUserService(mockRepo)
	t.Run("GetUserService Error", func(t *testing.T) {
		mockRepo.On("FindAll").Return([]models.User{}, errors.New("databse error"))

		user, err := service.GetUsersService()

		assert.Error(t, err)
		assert.Empty(t, user)
		mockRepo.AssertExpectations(t)
	})
}

func TestGetUserByIdServiceSuccess(t *testing.T) {
	t.Run("GetUserByIdService Success", func(t *testing.T) {
		expectedUser := &models.User{
			Id:    1,
			Name:  "Muskan",
			Email: "muskan@gmail.com",
			Age:   25,
		}

		mockRepo.On("FindById", 1).Return(expectedUser, nil)
		user, err := service.GetUserByIdService(1)

		assert.NoError(t, err)
		assert.Equal(t, expectedUser, user)
		mockRepo.AssertExpectations(t)
	})
}

func TestGetUserByIdServiceError(t *testing.T) {
	t.Run("GetUserByIdService Not Found", func(t *testing.T) {
		mockRepo.On("FindById", 999999).Return(nil, gorm.ErrRecordNotFound)
		user, err := service.GetUserByIdService(999999)

		assert.Error(t, err)
		assert.Nil(t, user)
		mockRepo.AssertExpectations(t)
	})
}

func TestGetUserByEmailServiceSuccess(t *testing.T) {
	t.Run("GetUserByEmailService Success", func(t *testing.T) {
		expectedUser := &models.User{
			Id:    1,
			Name:  "Muskan",
			Email: "muskan@gmail.com",
			Age:   25,
		}

		mockRepo.On("FindByEmail", "muskan@gmail.com").Return(expectedUser, nil)
		user, err := service.GetUserByEmailService("muskan@gmail.com")

		assert.NoError(t, err)
		assert.Equal(t, expectedUser, user)
		mockRepo.AssertExpectations(t)
	})
}

func TestGetUserByEmailServiceError(t *testing.T) {
	mockRepo = new(repositories.MockUserRepository)
	service = NewUserService(mockRepo)
	t.Run("GetUserByEmailService Not Found", func(t *testing.T) {
		mockRepo.On("FindByEmail", "muskan@gmail.com").Return(nil, gorm.ErrRecordNotFound)
		user, err := service.GetUserByEmailService("muskan@gmail.com")

		assert.Error(t, err)
		assert.Nil(t, user)
		mockRepo.AssertExpectations(t)
	})
}
