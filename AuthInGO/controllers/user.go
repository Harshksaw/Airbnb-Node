package controllers

import (
	"AuthInGO/services"
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

func (uc *UserController) CreateUser(w http.ResponseWriter , r *http.Request) error {
	// Call the user service to create a user
	uc.UserService.CreateUser()
	w.Write([]byte("User created successfully"))

}