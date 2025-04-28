package repos

import (
	"blog-api-go/internal/models"
	"database/sql"
	"fmt"
)

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
func (r RepositoryImpl) GetProfileUser(UserID int) (models.User, int, error) {
	var user models.User

	row := r.db.QueryRow(`SELECT id, login, email, password, full_name, photo FROM user_profiles WHERE id = $1`, UserID)

	err := row.Scan(&user.Id, &user.Login, &user.Email, &user.Password, &user.FullNameUser, &user.Photo)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, 0, fmt.Errorf("пользователь с id %d не найден", UserID)
		}
		return user, 0, err
	}

	return user, 1, nil
}
func (r RepositoryImpl) GetProfileUserForLogin(login string) ([]models.User, int, error) {
	var users []models.User

	searchQuery := "%" + login + "%"

	rows, err := r.db.Query(`SELECT id, login, email, password, full_name, photo FROM user_profiles WHERE login LIKE $1`, searchQuery)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.Login, &user.Email, &user.Password, &user.FullNameUser, &user.Photo); err != nil {
			return nil, 0, err
		}
		users = append(users, user)
	}

	if len(users) == 0 {
		return nil, 0, fmt.Errorf("Пользователи с логином, содержащим '%s', не найдены", login)
	}

	return users, 1, nil
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

func (r RepositoryImpl) LoginCheck(login string) bool {
	var exist bool

	err := r.db.QueryRow(`SELECT EXISTS (SELECT 1 FROM user_profiles WHERE login = $1)`, login).Scan(&exist)
	if err != nil {
		return false
	}
	return exist
}

func (r RepositoryImpl) SignUp(user models.User, hashPass models.Credentials) error {
	exists := r.LoginCheck(user.Login)
	if exists {
		return fmt.Errorf("логин уже существует")
	}

	_, err := r.db.Exec(`INSERT INTO user_profiles (login, email, password, full_name, photo) 
		VALUES ($1, $2, $3, $4, $5)`,
		user.Login, user.Email, hashPass.Password, user.FullNameUser, user.Photo)
	if err != nil {
		return fmt.Errorf("ошибка при регистрации: %v", err)
	}

	return nil
}
