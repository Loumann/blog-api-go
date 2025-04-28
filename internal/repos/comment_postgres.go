package repos

import (
	"blog-api-go/internal/models"
	"time"
)

func (r RepositoryImpl) CreateComment(userId int, postId int, comment models.Comments) (models.Comments, error) {
	var date = time.Now().Format("2006-01-02 15:04:05")

	row, err := r.db.Query(`insert into comments(user_id, content, date_create, id_post) 
	values ($1,$2,$3,$4)`,
		userId, comment.Content, date, postId)
	if row != nil {
		return comment, err
	}
	return comment, err
}
func (r RepositoryImpl) GetComments(idPost int) ([]models.Comments, error) {
	var commets []models.Comments
	rows, err := r.db.Query(`SELECT c.id_post, c.content, c.user_id,c.date_create, u.login, u.photo, u.full_name
	FROM comments c
         JOIN user_profiles u ON c.user_id = u.id
	WHERE c.id_post = $1 `, idPost)

	if err != nil {
		return commets, err
	}
	defer rows.Close()

	for rows.Next() {
		var comm models.Comments

		if err := rows.Scan(&comm.Id_post, &comm.Content, &comm.User_id, &comm.Date_created, &comm.Login, &comm.Photo, &comm.FullName); err != nil {
			return commets, err
		}
		commets = append(commets, comm)
	}

	if err := rows.Err(); err != nil {
		return commets, err
	}
	return commets, err
}
func (r RepositoryImpl) DeleteComment(CommentId int) error {
	_, err := r.db.Exec(`DELETE FROM comments WHERE id = $1`, CommentId)
	if err != nil {
		return err
	}
	return nil
}
func (r RepositoryImpl) ChangeComment(comment models.Comments) (bool, error) {
	var date = time.Now().Format("2006-01-02 15:04:05")
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
