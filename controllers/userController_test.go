package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"userRepo/models"
	"userRepo/services"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestGetUsersSuccess(t *testing.T) {
	services := new(services.MockUserService)

	t.Run("GetUsers Success", func(t *testing.T) {
		expectedUsers := []models.User{{
			Id:    1,
			Name:  "Muskan",
			Email: "muskan@gmail.com",
			Age:   25,
		},
		}
		services.On("GetUsersService").Return(expectedUsers, nil)
		controller := NewUserController(services)

		req, _ := http.NewRequest("GET", "/api/users", nil)
		w := httptest.NewRecorder()
		controller.GetUsers(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		services.AssertExpectations(t)
	})
}

func TestGetUsersError(t *testing.T) {
	services := new(services.MockUserService)

	t.Run("GetUsers Success", func(t *testing.T) {
		services.On("GetUsersService").Return([]models.User{}, assert.AnError)
		controller := NewUserController(services)
		req, _ := http.NewRequest("GET", "/api/users", nil)
		w := httptest.NewRecorder()

		controller.GetUsers(w, req)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

func TestGetUserByIdSuccess(t *testing.T) {
	services := new(services.MockUserService)

	t.Run("GetUserById Success", func(t *testing.T) {
		expectedUser := &models.User{
			Id:    1,
			Name:  "Muskan",
			Email: "muskan@gmail.com",
			Age:   25,
		}
		services.On("GetUserByIdService", 1).Return(expectedUser, nil)
		controller := NewUserController(services)
		req, _ := http.NewRequest("GET", "/api/user/1", nil)
		req = mux.SetURLVars(req, map[string]string{"id": "1"})
		w := httptest.NewRecorder()

		controller.GetUserById(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		services.AssertExpectations(t)
	})
}

func TestGetUserByIdError(t *testing.T) {
	services := new(services.MockUserService)

	t.Run("GetUserById Not Found", func(t *testing.T) {
		services.On("GetUserByIdService", 100).Return(nil, assert.AnError)

		controller := NewUserController(services)
		req, _ := http.NewRequest("GET", "/api/user/1", nil)
		req = mux.SetURLVars(req, map[string]string{"id": "100"})
		w := httptest.NewRecorder()

		controller.GetUserById(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		services.AssertExpectations(t)
	})
}

func TestGetUserByEmail(t *testing.T) {
	services := new(services.MockUserService)

	t.Run("GetUserByEmail Success", func(t *testing.T) {
		expectedUser := &models.User{
			Id:    1,
			Name:  "Muskan",
			Email: "muskan@gmail.com",
			Age:   25,
		}

		services.On("GetUserByEmailService", "muskan@gmail.com").Return(expectedUser, nil)
		controller := NewUserController(services)
		req, _ := http.NewRequest("GET", "/api/user/email?email=muskan@gmail.com", nil)
		w := httptest.NewRecorder()

		controller.GetUserByEmail(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		services.AssertExpectations(t)
	})
}

func TestGetUserByEmailError(t *testing.T) {
	services := new(services.MockUserService)

	t.Run("GetUserByEmail Not Found", func(t *testing.T) {
		services.On("GetUserByEmailService", "muskan@gmail.com").Return(nil, assert.AnError)

		controller := NewUserController(services)
		req, _ := http.NewRequest("GET", "/api/user/email?email=muskan@gmail.com", nil)
		w := httptest.NewRecorder()

		controller.GetUserByEmail(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		services.AssertExpectations(t)

	})
}
