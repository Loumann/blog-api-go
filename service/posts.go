package service

import "blog-api-go/models"

func (s *Services) GetPosts() (posts []models.Post, err error) {
	return s.Repository.GetPosts()
}
