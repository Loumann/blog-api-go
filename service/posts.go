package service

import "blog-api-go/models"

func (s *Services) GetComments() (posts []models.Comments, err error) {
	return s.Repository.GetComments()
}

func (s *Services) CreatePost(UserID int, post models.Post) error {
	return s.Repository.CreatePost(UserID, post)
}

func (s Services) CreateComment(userId int, postId int, comment models.Comments) (models.Comments, error) {
	return s.Repository.CreateComment(userId, postId, comment)
}

func (s Services) GetIdPost(userId int) (int, error) {
	return s.Repository.GetIdPost(userId)

}
func (s Services) GetPosts() ([]models.Post, error) {
	return s.Repository.GetPosts()
}
