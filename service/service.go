package service

import (
	"blog-api-go/models"
	"blog-api-go/repos"
)

type Service interface {
	SignIn(login string) ([]byte, int, error)
	SignUp(user models.User, hashPass models.Credentials) error

	GetProfileUserForLogin(login string) ([]models.User, int, error)
	GetAllUsers() ([]models.User, error)
	GetIdPost(userId int) (int, error)
	GetComments(idPost int) ([]models.Comments, error)
	GetProfileUser(UserID int) (models.User, int, error)

	ToggleSub(userID int, targetID string) bool
	CheckIfSubscribed(userID int, targetID string) (bool, error)

	GetPosts(userID, page, limit int, own bool) ([]models.Post, error)

	DeletePost(postId int) error
	DeleteComment(CommentId int) error

	CreateComment(userId int, postId int, comment models.Comments) (models.Comments, error)
	CreatePost(UserID int, post models.Post) error

	ChangePost(post models.Post) (bool, error)
	ChangeComment(comment models.Comments) (bool, error)
}

func NewService(Repos *repos.RepositoryImpl) *Services {
	return &Services{Repos}
}

type Services struct {
	Repository repos.Repository
}
