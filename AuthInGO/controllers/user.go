package controllers

import (
	"AuthInGo/services"
	"fmt"
	"net/http"
)


type UserController struct{
	UserService services.UserService
}

func NewUserCOntroller(_userService services.UserService) *UserController {
	return &UserController{
		UserService: _userService,
	}
}
func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
    err := uc.UserService.CreateUser()
    if err != nil {
        http.Error(w, "Failed to create user", http.StatusInternalServerError)
        return
    }
    w.Write([]byte("User created successfully"))
}

func (uc *UserController) GetUserById(w http.ResponseWriter , r *http.Request) error {
	// Call the user service to create a user
	fmt.Println("Getting user in UserController")
	uc.UserService.CreateUser();
	w.Write([]byte("User created successfully"))


	return nil

}