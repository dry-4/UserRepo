package repositories

import (
	"userRepo/models"

	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

var _ UserRepository = &MockUserRepository{}

func (m *MockUserRepository) FindAll() ([]models.User, error) {
	result := m.Called()
	return result.Get(0).([]models.User), result.Error(1)
}

func (m *MockUserRepository) FindById(id int) (*models.User, error) {
	result := m.Called(id)
	if result.Get(0) == nil {
		return nil, result.Error(1)
	}
	return result.Get(0).(*models.User), result.Error(1)
}

func (m *MockUserRepository) FindByEmail(email string) (*models.User, error) {
	result := m.Called(email)
	if result.Get(0) == nil {
		return nil, result.Error(1)
	}
	return result.Get(0).(*models.User), result.Error(1)
}
