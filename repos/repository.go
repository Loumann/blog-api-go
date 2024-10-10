package repos

import (
	"blog-api-go/models"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	SignIn(login string, password string) (err error)
	GetAllUsers() ([]models.User, error)
	GetPosts() ([]models.Post, error)
}

func NewRepository(db *sqlx.DB) *RepositoryImpl {
	return &RepositoryImpl{db: db}
}

type RepositoryImpl struct {
	db *sqlx.DB
}

func (r *RepositoryImpl) SignIn(login string, password string) error {
	query := `SELECT * FROM user_profiles WHERE login = $1 AND password = $2`
	var user models.User
	err := r.db.Get(&user, query, login, password)
	if err != nil {
		return err
	}
	return nil
}

func (r *RepositoryImpl) FindAll() ([]models.User, error) {
	var users []models.User
	rows, err := r.db.Query(`SELECT * FROM user`)
	if err != nil {
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.Login, &user.Email, &user.Password, &user.FullNameUser, &user.BirthDate, &user.Photo); err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}
func (r *RepositoryImpl) GetAllUsers() ([]models.User, error) {
	var users []models.User
	rows, err := r.db.Query(`SELECT * FROM user_profiles`)
	if err != nil {
		return users, err
	}
	defer rows.Close()
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.Login, &user.Email, &user.Password, &user.FullNameUser, &user.BirthDate, &user.Photo); err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *RepositoryImpl) GetPosts() ([]models.Post, error) {
	var posts []models.Post

	rows, err := r.db.Query(`SELECT * FROM posts`)
	if err != nil {
		return posts, err
	}
	defer rows.Close()

	for rows.Next() {
		var post models.Post

		if err := rows.Scan(&post.Id, &post.User_id, &post.Content, &post.Date_created, &post.Date_changed); err != nil {
			return posts, err
		}
		posts = append(posts, post)
	}
	if err := rows.Err(); err != nil {
		return posts, err
	}
	return posts, err
}
