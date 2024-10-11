package service

import (
	"blog-api-go/models"
	"blog-api-go/repos"
)

type Service interface {
	GetAllUsers() ([]models.User, error)
	SignIn(login string) ([]byte, error)
	GetPosts() ([]models.Post, error)
	SignUp(user models.User, hashPass models.Credentials) error
}

func NewService(Repos *repos.RepositoryImpl) *Services {
	return &Services{Repos}
}

type Services struct {
	Repository repos.Repository
}
