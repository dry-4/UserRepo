package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"userRepo/services"

	"github.com/gorilla/mux"
)

type UserController struct {
	service services.UserService
}

func NewUserController(service services.UserService) *UserController {
	return &UserController{service: service}
}

func (u *UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	users, err := u.service.GetUsersService()

	if err != nil {
		http.Error(w, "Error fetching users...", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func (u *UserController) GetUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {
		http.Error(w, "Invalid user ID...", http.StatusBadRequest)
		return
	}

	user, err := u.service.GetUserByIdService(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (u *UserController) GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	email := r.URL.Query().Get("email")
	fmt.Println(email)
	if email == "" {
		http.Error(w, "Invalid email...", http.StatusBadRequest)
		return
	}

	user, err := u.service.GetUserByEmailService(email)

	if err != nil {
		http.Error(w, "User not found", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)

}
