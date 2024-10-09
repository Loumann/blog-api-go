package service

import (
	"blog-api-go/models"
	"blog-api-go/repos"
)

type UserService interface {
	GetAllUsers() ([]models.User, error)
	SignIn(login string, password string) error
	FindAll() ([]models.User, error)
}

func NewUserService(userRepo UserService) UserService {
	return &userService{userRepo}
}

type userService struct {
	userRepository repos.UserRepository
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.userRepository.GetAllUsers()
}

func (s *userService) SignIn(login string, password string) error {
	return s.userRepository.SignIn(login, password) // Передаем логин и пароль
}
func (s *userService) FindAll() ([]models.User, error) {
	return s.userRepository.FindAll()
}
