package service

import (
	models2 "blog-api-go/internal/models"
	"blog-api-go/internal/repos"
)

type User interface {
	SignIn(login string) ([]byte, int, error)
	SignUp(user models2.User, hashPass models2.Credentials) error

	GetProfileUserForLogin(login string) ([]models2.User, int, error)
	GetAllUsers() ([]models2.User, error)
	GetProfileUser(UserID int) (models2.User, int, error)
}

type Subscribe interface {
	ToggleSub(userID int, targetID string) bool
	CheckIfSubscribed(userID int, targetID string) (bool, error)
}

type Comment interface {
	GetComments(idPost int) ([]models2.Comments, error)
	CreateComment(userId int, postId int, comment models2.Comments) (models2.Comments, error)
	ChangeComment(comment models2.Comments) (bool, error)
	DeleteComment(CommentId int) error
}
type Post interface {
	GetIdPost(userId int) (int, error)
	GetPosts(userID, page, limit int, own bool) ([]models2.Post, error)
	CreatePost(UserID int, post models2.Post) error
	ChangePost(post models2.Post) (bool, error)
	DeletePost(postId int) error
}

type Service struct {
	Post
	User
	Comment
	Subscribe
}

func NewService(Repos *repos.Repository) *Services {
	return &Services{Repos}
}

type Services struct {
	Repository *repos.Repository
}
