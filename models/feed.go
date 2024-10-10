package models

import "time"

type Post struct {
	Id           int       `json:"id" db:"id"`
	User_id      int       `json:"user_id" db:"user_id"`
	Content      string    `json:"content" db:"content"`
	Date_created time.Time `json:"date_created" db:"date_created"`
	Date_changed time.Time `json:"date_updated" db:"date_changed"`
}
type Data time.Time

func (d Data) MarshalJSON() ([]byte, error) {
	formatted := "\"" + time.Time(d).Format("2006-01-02") + "\""
	return []byte(formatted), nil
}
