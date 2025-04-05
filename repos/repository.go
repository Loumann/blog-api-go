package repos

import (
	"blog-api-go/models"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

var date = time.Now().Format("2006-01-02 15:04:05")

type RepositoryImpl struct {
	db *sqlx.DB
}
type Repository interface {
	SignIn(login string) ([]byte, int, error)
	SignUp(user models.User, hashPass models.Credentials) (err error)

	GetProfileUserForLogin(login string) ([]models.User, int, error)

	GetAllUsers() ([]models.User, error)
	GetProfileUser(UserID int) (models.User, int, error)
	GetComments() ([]models.Comments, error)
	GetIdPost(PostId int) (int, error)
	GetPosts(userID, page, limit int, own bool) ([]models.Post, error)

	Subscribe(UserId string, Subscriber int) (bool, error)
	IsSubscribe(UserId string, Subscriber int) (error, bool)

	CreatePost(UserID int, post models.Post) error
	CreateComment(userId int, postId int, comment models.Comments) (models.Comments, error)

	DeletePost(PostId int) error
	DeleteComment(CommentId int) error

	ChangePost(post models.Post) (bool, error)
	ChangeComment(comments models.Comments) (bool, error)
}

func NewRepository(db *sqlx.DB) *RepositoryImpl {
	return &RepositoryImpl{db: db}
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

	// Проверяем, есть ли найденные пользователи
	if len(users) == 0 {
		return nil, 0, fmt.Errorf("Пользователи с логином, содержащим '%s', не найдены", login)
	}

	return users, 1, nil
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

		if err := rows.Scan(&comm.Id, &comm.User_id, &comm.Content, &comm.Date_created, &comm.Id_post_on_comment); err != nil {
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
			return 0, fmt.Errorf("Пост с id %d не найден", postId)
		}
		return 0, err
	}
	return idPost, nil
}
func (r RepositoryImpl) GetPosts(userID, page, limit int, own bool) ([]models.Post, error) {
	var posts []models.Post
	offset := (page - 1) * limit

	var rows *sql.Rows
	var err error

	if own == true {
		rows, err = r.db.Query(`SELECT p.id_post, p.id_user_create_post, p.theme, p.content_post, p.date_create, 
                   u.full_name, u.photo, u.login, LENGTH(p.content_post) > 500 as is_long
            FROM post p 
            JOIN user_profiles u ON p.id_user_create_post = u.id 
            WHERE p.id_user_create_post = $1 
            ORDER BY date_create DESC
            LIMIT $2 OFFSET $3`, userID, limit, offset)
	} else {
		rows, err = r.db.Query(`
            SELECT p.id_post, p.id_user_create_post, p.theme, p.content_post, p.date_create, 
                   u.full_name, u.photo, u.login, LENGTH(p.content_post) > 500 as is_long
            FROM post p 
            JOIN user_profiles u ON p.id_user_create_post = u.id 
            ORDER BY date_create DESC
            LIMIT $1 OFFSET $2`, limit, offset)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post models.Post
		var isLong bool

		if err := rows.Scan(
			&post.Id_post,
			&post.Id_User,
			&post.Theme,
			&post.Content_post,
			&post.Date_create,
			&post.FullName,
			&post.Photo,
			&post.Login,
			&isLong,
		); err != nil {
			return nil, err
		}

		post.IsLong = isLong
		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func (r RepositoryImpl) CreateComment(userId int, postId int, comment models.Comments) (models.Comments, error) {

	row, err := r.db.Query(`insert into comments(user_id, content, date_created, id_post_on_comment) 
	values ($1,$2,$3,$4)`,
		userId, comment.Content, date, postId)
	if row != nil {
		return comment, err
	}
	return comment, err
}
func (r RepositoryImpl) CreatePost(UserID int, post models.Post) error {

	row, err := r.db.Query(`INSERT INTO post (id_user_create_post, theme, content_post, date_create) 
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

	row, err := r.db.Query(`INSERT INTO user_profiles (login, email, password, full_name,  photo) 
	VALUES ($1, $2,$3,$4,$5)`,
		user.Login, user.Email, hashPass.Password, user.FullNameUser, user.Photo)
	if row != nil {
		return err
	}
	return err
}

func (r RepositoryImpl) DeletePost(postId int) error {

	_, err := r.db.Exec(`DELETE FROM comments WHERE id_post_on_comment = $1`, postId)
	if err != nil {
		return err
	}

	_, err = r.db.Exec(`DELETE FROM post WHERE id_post = $1`, postId)
	if err != nil {
		return err
	}

	return nil
}
func (r RepositoryImpl) DeleteComment(CommentId int) error {
	_, err := r.db.Exec(`DELETE FROM comments WHERE id = $1`, CommentId)
	if err != nil {
		return err
	}
	return nil
}

func (r RepositoryImpl) ChangePost(post models.Post) (bool, error) {

	var count int
	err := r.db.QueryRow(`select count(*) from post where id_post=$1`, post.Id_post).Scan(&count)
	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, nil
	}

	_, err = r.db.Exec(`UPDATE post SET theme=$1, content_post=$2, date_create=$3 WHERE id_post=$4`,
		post.Theme, post.Content_post, date, post.Id_post)
	return true, err
}
func (r RepositoryImpl) ChangeComment(comment models.Comments) (bool, error) {
	var count int

	err := r.db.QueryRow(`SELECT count(*) from comments where id=$1`, comment.Id).Scan(&count)
	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, nil
	}
	_, err = r.db.Exec(`UPDATE comment set content=$1, data_created=$2`, comment.Content, date)
	if err != nil {
		return false, err
	}
	return true, nil

}

func (r RepositoryImpl) Subscribe(UserId string, Subscriber int) (bool, error) {
	var exist bool

	qr := r.db.QueryRow(`SELECT exists(SELECT 1 FROM subscribe where subscriber_id=$2 AND subscribed_to_id=$2)`, UserId, Subscriber)
	if err := qr.Err(); err != nil {
		return false, err
	}
	if exist {
		return false, nil
	}
	_, err := r.db.Exec(`INSERT into subscribe (subscriber_id,subscribed_to_id, date_subscribed )
		values ($1, $2, $3)`,
		UserId, Subscriber, date)

	return err == nil, err

}
func (r RepositoryImpl) UnSubscribe(UserId string, Subscriber int) error {
	qr, err := r.db.Query(`DELETE FROM subscribe WHERE subscriber_id=$1 and subscriber_id=$2`, UserId, Subscriber)
	if err != nil {
		return err
	}
	print(qr)
	return err
}

func (r RepositoryImpl) IsSubscribe(UserId string, Subscriber int) (error, bool) {
	qr := r.db.QueryRow(`SELECT EXISTS(SELECT 1 FROM subscribe WHERE subscriber_id = $1 AND subscribed_to_id = $2)`, Subscriber, UserId)
	var exists bool
	err := qr.Scan(&exists)
	if err != nil {
		return err, false
	}
	println(exists)
	return err, exists
}
