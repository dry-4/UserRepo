package main

import (
	"fmt"
	"log"
	"net/http"
	"userRepo/controllers"
	"userRepo/repositories"
	"userRepo/services"

	"github.com/gorilla/mux"
)

func main() {

	userRepository := repositories.NewUserRepository()
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	router := mux.NewRouter()
	router.HandleFunc("/api/users", userController.GetUsers).Methods("GET")
	router.HandleFunc("/api/users/{id}", userController.GetUserById).Methods("GET")
	router.HandleFunc("/api/user/email", userController.GetUserByEmail).Methods("GET")

	fmt.Println("Staring server on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}
