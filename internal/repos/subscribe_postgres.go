package repos

import "log"

func (r RepositoryImpl) CheckIfSubscribed(userID int, targetID string) (bool, error) {
	var exists bool

	err := r.db.QueryRow(`SELECT EXISTS (
        SELECT 1 FROM subscribe WHERE subscriber_id=$1 AND subscribed_to_id=$2
    )`, userID, targetID).Scan(&exists)
	if err != nil {
		log.Println("Ошибка при проверке подписки:", err)
		return false, err
	}
	return exists, err
}
func (r RepositoryImpl) ToggleSub(userID int, targetID string) bool {
	subscribed, err := r.CheckIfSubscribed(userID, targetID)
	if err != nil {
		return false
	}

	if subscribed {
		_, err := r.db.Exec(`DELETE FROM subscribe WHERE subscriber_id=$1 AND subscribed_to_id=$2`, userID, targetID)
		if err != nil {
			log.Println("Ошибка при отписке:", err)
		}
		return false
	} else {
		_, err := r.db.Exec(`INSERT INTO subscribe (subscriber_id, subscribed_to_id) VALUES ($1, $2)`, userID, targetID)
		if err != nil {
			log.Println("Ошибка при подписке:", err)
		}
		return true
	}
}
