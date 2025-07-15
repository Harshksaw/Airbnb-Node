package services

import (
	db "AuthInGO/db/repositories"
	"fmt"
)

type UserService interface{
	CreateUser() error
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
	
	return nil
}