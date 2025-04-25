package service

func (s *Services) ToggleSub(userID int, targetID string) bool {
	return s.Repository.ToggleSub(userID, targetID)
}

func (s *Services) CheckIfSubscribed(userID int, targetID string) (bool, error) {
	return s.Repository.CheckIfSubscribed(userID, targetID)
}
