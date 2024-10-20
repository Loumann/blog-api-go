package models

import "time"

type Comments struct {
	Id                 int       `json:"id_comment" db:"id"`
	User_id            int       `json:"user_id" db:"user_id"`
	Content            string    `json:"content" db:"content"`
	Date_created       time.Time `json:"date_created" db:"date_created"`
	Id_post_on_comment int       `json:"id_post_on_comment" db:"id_post_on_comment"`
}

type Post struct {
	Id_post      int    `json:"id_post" db:"id_post"`
	Id_User      int    `json:"id_user_create_post" db:"id_user_create_post"`
	Theme        string `json:"theme" db:"theme"`
	Content_post string `json:"content_post" db:"content_post"`
	Date_create  string `json:"date_create_post" db:"date_create_post"`
}
