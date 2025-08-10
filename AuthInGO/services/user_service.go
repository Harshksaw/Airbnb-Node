package services

import (
	db "AuthInGo/db/repositories"
	"fmt"


)

type UserService interface{
	CreateUser() error
}
type UserService interface {
	GetUserByID(id int) (User, error)
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

func (u *UserServiceImpl) GetUserById() error {
	// Implementation for getting a user by ID
	fmt.Println("Fetching user in UserService")
	// u.userRepository.GetById()
	return nil
}
