package service

import (
	"blog-api-go/models"
	"blog-api-go/repos"
)

type Service interface {
	SignIn(login string) ([]byte, int, error)
	SignUp(user models.User, hashPass models.Credentials) error

	GetAllUsers() ([]models.User, error)
	GetIdPost(userId int) (int, error)
	GetComments() ([]models.Comments, error)
	GetProfileUser(UserID int) ([]models.User, int, error)
	GetPosts() ([]models.Post, error)

	CreateComment(userId int, postId int, comment models.Comments) (models.Comments, error)
	CreatePost(UserID int, post models.Post) error
}

func NewService(Repos *repos.RepositoryImpl) *Services {
	return &Services{Repos}
}

type Services struct {
	Repository repos.Repository
}
