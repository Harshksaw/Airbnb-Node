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
	// payload := r.Context().Value("payload").(dto.CreateUserRequestDTO)

	fmt.Println("Payload received:")

	user, err := uc.UserService.CreateUser()

	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Failed to create user", err)
		return
	}

	// utils.WriteJsonSuccessResponse(w, http.StatusCreated, "User created successfully", user)
	fmt.Println("User created successfully:", user)
}

func (uc *UserController) GetUserById(w http.ResponseWriter , r *http.Request) error {
	// Call the user service to create a user
	fmt.Println("Getting user in UserController")
	// uc.UserService.CreateUser();
	w.Write([]byte("User created successfully"))


	return nil

}
