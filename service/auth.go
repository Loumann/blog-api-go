package service

import "blog-api-go/models"

func (s *Services) GetAllUsers() ([]models.User, error) {
	return s.Repository.GetAllUsers()
}

func (s *Services) SignIn(login string, password string) error {
	return s.Repository.SignIn(login, password)
}
