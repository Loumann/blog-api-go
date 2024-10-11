package service

import "blog-api-go/models"

func (s *Services) GetAllUsers() ([]models.User, error) {
	return s.Repository.GetAllUsers()
}

func (s *Services) SignIn(login string) ([]byte, error) {
	return s.Repository.SignIn(login)
}

func (s *Services) SignUp(user models.User, pass models.Credentials) error {
	return s.Repository.SignUp(user, pass)
}
