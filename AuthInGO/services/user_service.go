package services

import (

	db "AuthInGo/db/repositories"
	"fmt"
	"AuthInGo/models"


)


type UserService interface {
	GetUserByID(id string) (*models.User, error)
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