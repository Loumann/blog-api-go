package repos

import (
	"blog-api-go/models"
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	FindAll() ([]models.User, error)
	SignIn(login string, password string) (err error)
	GetAllUsers() ([]models.User, error)
}

func NewUserRepository(db *sqlx.DB) *userRepositoryImpl {
	return &userRepositoryImpl{db: db}
}

type userRepositoryImpl struct {
	db *sqlx.DB
}

func (r *userRepositoryImpl) FindAll() ([]models.User, error) {
	var users []models.User
	rows, err := r.db.Query(`SELECT * FROM login`)
	if err != nil {
		return users, err
	}
	defer rows.Close()
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.Login, &user.Email, &user.Password, &user.IdProfile); err != nil {
			return users, err
		}
		users = append(users, user)

	}

	return users, nil
}
func (r *userRepositoryImpl) SignIn(login string, password string) error {
	query := `SELECT * FROM user WHERE login = $1 AND password = $2`
	var user models.User
	err := r.db.Get(&user, query, login, password)
	if err != nil {
		return err
	}
	return nil
}
func (r *userRepositoryImpl) GetAllUsers() ([]models.User, error) {

	var users []models.User
	rows, err := r.db.Query(`SELECT * FROM login`)
	if err != nil {
		return users, err
	}
	defer rows.Close()
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.Login, &user.Email, &user.Password, &user.IdProfile); err != nil {
			return users, err
		}
		users = append(users, user)

	}

	return users, nil

}
