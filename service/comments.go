package service

import "blog-api-go/models"

func (s Services) CreateComment(userId int, postId int, comment models.Comments) (models.Comments, error) {
	return s.Repository.CreateComment(userId, postId, comment)
}

func (s Services) GetComments() (posts []models.Comments, err error) {
	return s.Repository.GetComments()
}

func (s Services) DeleteComment(CommentId int) error {
	return s.Repository.DeleteComment(CommentId)
}

func (s Services) ChangeComment(commentId models.Comments) (bool, error) {
	return s.Repository.ChangeComment(commentId)
}
