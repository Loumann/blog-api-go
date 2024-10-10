package service

import (
	"blog-api-go/models"
	"blog-api-go/repos"
)

type Service interface {
	GetAllUsers() ([]models.User, error)
	SignIn(login string, password string) error
	GetPosts() ([]models.Post, error)
}

func NewService(Repos *repos.RepositoryImpl) *Services {
	return &Services{Repos}
}

type Services struct {
	Repository repos.Repository
}
