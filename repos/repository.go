package repos

import (
	"blog-api-go/models"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

type Repository interface {
	SignIn(login string) ([]byte, int, error)
	SignUp(user models.User, hashPass models.Credentials) (err error)

	GetAllUsers() ([]models.User, error)
	GetProfileUser(UserID int) ([]models.User, int, error)
	GetComments() ([]models.Comments, error)
	GetIdPost(PostId int) (int, error)
	GetPosts() ([]models.Post, error)

	CreatePost(UserID int, post models.Post) error
	CreateComment(userId int, postId int, comment models.Comments) (models.Comments, error)

	DeletePost(PostId int) error
}

func NewRepository(db *sqlx.DB) *RepositoryImpl {
	return &RepositoryImpl{db: db}
}

type RepositoryImpl struct {
	db *sqlx.DB
}

func (r RepositoryImpl) GetAllUsers() ([]models.User, error) {
	var users []models.User
	rows, err := r.db.Query(`SELECT * FROM user_profiles`)
	if err != nil {
		return users, err
	}
	defer rows.Close()
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.Login, &user.Email, &user.Password, &user.FullNameUser, &user.Photo); err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}
func (r RepositoryImpl) GetProfileUser(UserID int) ([]models.User, int, error) {
	var users []models.User

	rows, err := r.db.Query(`SELECT * FROM user_profiles WHERE id = $1`, UserID)
	if err != nil {

		return users, 0, err
	}
	defer rows.Close()
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.Login, &user.Email, &user.Password, &user.FullNameUser, &user.Photo); err != nil {
			return users, 0, err
		}
		users = append(users, user)
	}
	return users, 0, nil
}
func (r RepositoryImpl) GetComments() ([]models.Comments, error) {
	var commets []models.Comments
	rows, err := r.db.Query(`SELECT * FROM comments`)
	if err != nil {
		return commets, err
	}
	defer rows.Close()
	for rows.Next() {
		var comm models.Comments

		if err := rows.Scan(&comm.Id, &comm.User_id, &comm.Content, &comm.Date_created); err != nil {
			return commets, err
		}
		commets = append(commets, comm)
	}
	if err := rows.Err(); err != nil {
		return commets, err
	}
	return commets, err
}
func (r RepositoryImpl) GetIdPost(postId int) (int, error) {
	var idPost int
	err := r.db.QueryRow(`SELECT id_post FROM post WHERE id_post = $1`, postId).Scan(&idPost)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("пост с id %d не найден", postId)
		}
		return 0, err // Возвращаем ошибку, если возникла другая ошибка
	}
	return idPost, nil
}

func (r RepositoryImpl) CreateComment(userId int, postId int, comment models.Comments) (models.Comments, error) {
	date := time.Now().Format("2006-01-02 15:04:05")

	row, err := r.db.Query(`insert into comments(user_id, content, date_created, id_post_on_comment) 
	values ($1,$2,$3,$4)`,
		userId, comment.Content, date, postId)
	if row != nil {
		return comment, err
	}
	return comment, err
}
func (r RepositoryImpl) CreatePost(UserID int, post models.Post) error {
	date := time.Now().Format("2006-01-02 15:04:05")

	row, err := r.db.Query(`INSERT INTO post (id_user_create_post, theme, content_post, date_create_post) 
	VALUES ($1, $2,$3,$4)`,
		UserID, post.Theme, post.Content_post, date)
	if row != nil {
		return err
	}
	return err
}

func (r RepositoryImpl) SignIn(login string) ([]byte, int, error) {
	var hashedPassword string
	var userID int

	query := `SELECT password, id FROM user_profiles WHERE login = $1`

	err := r.db.QueryRow(query, login).Scan(&hashedPassword, &userID)
	if err != nil {
		return nil, 0, err
	}

	return []byte(hashedPassword), userID, nil
}
func (r RepositoryImpl) SignUp(user models.User, hashPass models.Credentials) error {

	row, err := r.db.Query(`INSERT INTO user_profiles (login, email, password, full_name_user,  photo) 
	VALUES ($1, $2,$3,$4,$5)`,
		user.Login, user.Email, hashPass.Password, user.FullNameUser, user.Photo)
	if row != nil {
		return err
	}
	return err
}

func (r RepositoryImpl) DeletePost(postId int) error {
	row, err := r.db.Query(`DELETE FROM post WHERE id_post = $1`, postId)
	if row != nil {
		return err
	}
	return err
}

func (r RepositoryImpl) GetPosts() ([]models.Post, error) {
	var posts []models.Post
	rows, err := r.db.Query(`SELECT * FROM post`)
	if err != nil {
		return posts, err
	}
	defer rows.Close()
	for rows.Next() {
		var post models.Post

		if err := rows.Scan(&post.Id_post, &post.Id_User, &post.Theme, &post.Content_post, &post.Date_create); err != nil {
			return posts, err
		}
		posts = append(posts, post)
	}
	if err := rows.Err(); err != nil {
		return posts, err
	}
	return posts, err
}
