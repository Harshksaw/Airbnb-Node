package services

import (
	db "AuthInGo/db/repositories"
	"AuthInGo/dto"
	"AuthInGo/models"
	"AuthInGo/utils"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type UserService interface {
	GetUserByID(id string) (*models.User, error)
	CreateUser(payload *dto.CreateUserRequestDTO) (*dto.CreateUserResponseDTO, error)
	LoginUser(payload *dto.LoginUserRequestDTO) (string, error)
}
type UserServiceImpl struct {
	userRepository db.UserRepository
}

func NewUserService(_userRepository db.UserRepository) UserService {
	return &UserServiceImpl{
		userRepository: _userRepository,
	}
}
func (u *UserServiceImpl) CreateUser(payload *dto.CreateUserRequestDTO) (*dto.CreateUserResponseDTO, error) {
	// Implementation for creating a user
	fmt.Println("Creating user in UserService")

	hashedPassword, err := utils.HashPassword(payload.Password)

	if err != nil {
		return nil, err
	}

	err = u.userRepository.Create(payload.Username, payload.Email, hashedPassword)
	if err != nil {
		return nil, err
	}

	return &dto.CreateUserResponseDTO{
		Id:       1, // This should ideally be returned from the database
		Username: payload.Username,
		Email:    payload.Email,
	}, nil
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

func (u *UserServiceImpl) LoginUser(payload *dto.LoginUserRequestDTO) (string, error) {
	email := payload.Email
	password := payload.Password

	user, err := u.userRepository.GetByEmail(email)
	if err != nil {
		return "", err
	}

	if user == nil {
		return "", fmt.Errorf("user not found")
	}

	isPasswordValid := utils.CheckPasswordHash(password, user.Password)
	if !isPasswordValid {
		return "", fmt.Errorf("invalid credentials")
	}

	jwtPayload := jwt.MapClaims{
		"email": user.Email,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtPayload)
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "TOKEN"
	}
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
