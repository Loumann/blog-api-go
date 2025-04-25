package repos

import (
	models2 "blog-api-go/internal/models"
	"github.com/jmoiron/sqlx"
	"time"
)

var date = time.Now().Format("2006-01-02 15:04:05")

type RepositoryImpl struct {
	db *sqlx.DB
}

type User interface {
	SignIn(login string) ([]byte, int, error)
	SignUp(user models2.User, hashPass models2.Credentials) (err error)

	GetProfileUserForLogin(login string) ([]models2.User, int, error)
	GetAllUsers() ([]models2.User, error)
	GetProfileUser(UserID int) (models2.User, int, error)
}
type Comment interface {
	GetComments(idPost int) ([]models2.Comments, error)
	CreateComment(userId int, postId int, comment models2.Comments) (models2.Comments, error)
	DeleteComment(CommentId int) error
	ChangeComment(comments models2.Comments) (bool, error)
}
type Post interface {
	GetIdPost(PostId int) (int, error)
	GetPosts(userID, page, limit int, own bool) ([]models2.Post, error)
	CreatePost(UserID int, post models2.Post) error
	DeletePost(PostId int) error
	ChangePost(post models2.Post) (bool, error)
}
type Subscribe interface {
	ToggleSub(userID int, targetID string) bool
	CheckIfSubscribed(userID int, targetID string) (bool, error)
}

type Repository struct {
	User
	Post
	Subscribe
	Comment
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User:      &RepositoryImpl{db: db},
		Post:      &RepositoryImpl{db: db},
		Subscribe: &RepositoryImpl{db: db},
		Comment:   &RepositoryImpl{db: db},
	}
}
