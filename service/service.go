package service

import (
	"blog-api-go/models"
	"blog-api-go/repos"
)

type Service interface {
	SignIn(login string) ([]byte, int, error)
	SignUp(user models.User, hashPass models.Credentials) error

	GetAllUsers() ([]models.User, error)
	GetPosts() ([]models.Post, error)
	GetProfileUser(UserID int) ([]models.User, int, error)
}

func NewService(Repos *repos.RepositoryImpl) *Services {
	return &Services{Repos}
}

type Services struct {
	Repository repos.Repository
}
