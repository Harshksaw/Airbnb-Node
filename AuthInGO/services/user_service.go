package services

import (
	db "AuthInGo/db/repositories"
	"AuthInGo/models"
	"AuthInGo/utils"
	"fmt"
)

type UserService interface {
	GetUserByID(id string) (*models.User, error)
	CreateUser() error
	LoginUser(email, password string) (*models.User, error)
}
type UserServiceImpl struct {
	userRepository db.UserRepository
}

func NewUserService(_userRepository db.UserRepository) UserService {
	return &UserServiceImpl{
		userRepository: _userRepository,
	}
}
func (u *UserServiceImpl) CreateUser() error {
	// Implementation for creating a user
	fmt.Println("Creating user in UserService")
	// u.userRepository.GetUserByID()

	password := "123456"
	hashedPassword := utils.HashPassword(password)

	if err !=  nil {
		return err
	}

	u.userRepository.Create("harsh", "user@gmamil.com", hashedPassword)

	return nil
}

func (u *UserServiceImpl) GetUserByID(id string) (*models.User, error) {
	fmt.Println("Fetching user in UserService")

	// Call the repository to fetch the user
	user, err := u.userRepository.GetById(id)
	if err != nil {
		fmt.Println("Error fetching user:", err)
		return nil, err
	}

	return user, nil
}


func ( u *UserServiceImpl) LoginUser(email, password string) (*models.User, error) {

	user, err := u.userRepository.GetByEmail(email)
	if err != nil {
		return nil, err
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, fmt.Errorf("invalid credentials")
	}
	//return JWT 

	

	utils.CheckPasswordHash(password,userRepository.GetByEmail(email).Password)


}