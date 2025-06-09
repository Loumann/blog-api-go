package service

import "blog-api-go/internal/models"

func (s *Services) ToggleSub(userID int, targetID string) bool {
	return s.Repository.ToggleSub(userID, targetID)
}

func (s *Services) CheckIfSubscribed(userID int, targetID string) (bool, error) {
	return s.Repository.CheckIfSubscribed(userID, targetID)
}

func (s *Services) GetPostFromSub(userID, page, limit int) ([]models.Post, error) {
	return s.Repository.GetSubscribedPosts(userID, page, limit)

}
