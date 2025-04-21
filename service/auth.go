package service

import "blog-api-go/models"

func (s *Services) GetAllUsers() ([]models.User, error) {
	return s.Repository.GetAllUsers()
}

func (s *Services) SignIn(login string) ([]byte, int, error) {
	return s.Repository.SignIn(login)
}

func (s *Services) SignUp(user models.User, pass models.Credentials) error {
	return s.Repository.SignUp(user, pass)
}

func (s *Services) GetProfileUser(UserId int) (models.User, int, error) {
	return s.Repository.GetProfileUser(UserId)
}
func (s *Services) GetProfileUserForLogin(login string) ([]models.User, int, error) {
	return s.Repository.GetProfileUserForLogin(login)

}

func (s *Services) ToggleSub(userID int, targetID string) bool {
	return s.Repository.ToggleSub(userID, targetID)
}

func (s *Services) CheckIfSubscribed(userID int, targetID string) bool {
	return s.Repository.CheckIfSubscribed(userID, targetID)
}
