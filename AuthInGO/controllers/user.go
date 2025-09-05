package controllers

import (
	"AuthInGo/dto"
	"AuthInGo/services"
	"AuthInGo/utils"
	"fmt"
	"net/http"
)

type UserController struct {
	UserService services.UserService
}

func NewUserController(_userService services.UserService) *UserController {
	return &UserController{
		UserService: _userService,
	}
}
func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	// payload := r.Context().Value("payload").(dto.CreateUserRequestDTO)

	fmt.Println("Payload received:")

	user, err := uc.UserService.CreateUser()

	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create user: %v", err), http.StatusInternalServerError)
		return
	}

	// utils.WriteJsonSuccessResponse(w, http.StatusCreated, "User created successfully", user)
	fmt.Println("User created successfully:", user)
}

func (uc *UserController) GetUserById(w http.ResponseWriter, r *http.Request) error {
	// Call the user service to create a user
	fmt.Println("Getting user in UserController")
	// uc.UserService.CreateUser();
	w.Write([]byte("User created successfully"))

	return nil

}

func (uc *UserController) LoginUser(w http.ResponseWriter, r *http.Request) {

	// payload := r.Context().Value("payload").(dto.LoginUserRequestDTO)

	var payload dto.LoginUserRequestDTO

	if jsonErr := utils.ReadJsonBody(r, &payload); jsonErr != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, jsonErr, "Invalid JSON body")
		return

	}

	if validationErr := utils.Validator.Struct(payload); validationErr != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, validationErr, "Invalid request payload")
		return
	}

	jwtToken, err := uc.UserService.LoginUser(&payload)


	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, err, "Failed to login user")	
		return
	}




	utils.WriteJsonSuccessResponse(w, http.StatusOK, "User created successfully",jwtToken)

}
