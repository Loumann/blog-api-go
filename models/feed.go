package models

import (
	"database/sql"
	"time"
)

type Comments struct {
	Id                 int       `json:"id_comment" db:"id"`
	User_id            int       `json:"user_id" db:"user_id"`
	Content            string    `json:"content" db:"content"`
	Date_created       time.Time `json:"date_created" db:"date_created"`
	Id_post_on_comment int       `json:"id_post_on_comment" db:"id_post_on_comment"`
}

type Post struct {
	Id_post      int64          `json:"id_post" db:"id_post"`
	Id_User      int64          `json:"id_user_create_post" db:"id_user_create_post"`
	Theme        string         `json:"theme" db:"theme"`
	Content_post string         `json:"content_post" db:"content_post"`
	Date_create  string         `json:"date_create" db:"date_create"`
	FullName     string         `json:"fullname" db:"fullname"`
	Photo        sql.NullString `json:"photo" db:"photo"`
	IsLong       bool           `json:"is_long" db:"is_long"`
	Login        string         `json:"login" db:"login"`
}
