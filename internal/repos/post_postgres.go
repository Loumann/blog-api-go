package repos

import (
	"blog-api-go/internal/models"
	"database/sql"
	"fmt"
)

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

func (r RepositoryImpl) CreatePost(UserID int, post models.Post) error {

	row, err := r.db.Query(`INSERT INTO post (id_user_create_post, theme, content_post, date_create) 
	VALUES ($1, $2,$3,$4)`,
		UserID, post.Theme, post.Content_post, date)
	if row != nil {
		return err
	}
	return err
}

func (r RepositoryImpl) DeletePost(postId int) error {

	_, err := r.db.Exec(`DELETE FROM comments WHERE id_post = $1`, postId)
	if err != nil {
		return err
	}

	_, err = r.db.Exec(`DELETE FROM post WHERE id_post = $1`, postId)
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
