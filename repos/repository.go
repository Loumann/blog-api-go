package repos

import (
	"blog-api-go/models"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	SignIn(login string) ([]byte, int, error)
	SignUp(user models.User, hashPass models.Credentials) (err error)

	GetAllUsers() ([]models.User, error)
	GetPosts() ([]models.Post, error)
	GetProfileUser(UserID int) ([]models.User, int, error)
}

func NewRepository(db *sqlx.DB) *RepositoryImpl {
	return &RepositoryImpl{db: db}
}

type RepositoryImpl struct {
	db *sqlx.DB
}

func (r RepositoryImpl) FindAll() ([]models.User, error) {
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
func (r RepositoryImpl) GetAllUsers() ([]models.User, error) {
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
func (r RepositoryImpl) GetPosts() ([]models.Post, error) {
	var posts []models.Post
	rows, err := r.db.Query(`SELECT * FROM comments`)
	if err != nil {
		return posts, err
	}
	defer rows.Close()
	for rows.Next() {
		var post models.Post

		if err := rows.Scan(&post.Id, &post.User_id, &post.Content, &post.Date_created); err != nil {
			return posts, err
		}
		posts = append(posts, post)
	}
	if err := rows.Err(); err != nil {
		return posts, err
	}
	return posts, err
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

	row, err := r.db.Query(`INSERT INTO user_profiles (login, email, password, full_name_user, birthday_user, photo) 
	VALUES ($1, $2,$3,$4,$5,$6)`,
		user.Login, user.Email, hashPass.Password, user.FullNameUser, user.BirthDate, user.Photo)
	if row != nil {
		return err
	}
	return err
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
