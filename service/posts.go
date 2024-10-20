package service

import "blog-api-go/models"

func (s *Services) CreatePost(UserID int, post models.Post) error {
	return s.Repository.CreatePost(UserID, post)
}

func (s Services) GetIdPost(userId int) (int, error) {
	return s.Repository.GetIdPost(userId)
}

func (s Services) GetPosts() ([]models.Post, error) {
	return s.Repository.GetPosts()
}

func (s Services) DeletePost(postId int) error {
	return s.Repository.DeletePost(postId)
}

func (s Services) ChangePost(post models.Post) (bool, error) {
	return s.Repository.ChangePost(post)
}
