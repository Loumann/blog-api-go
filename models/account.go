package models

type User struct {
	Id        int    `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Login     string `json:"login" db:"login"`
	Email     string `json:"email" db:"email"`
	Password  string `json:"password" db:"password"`
	IdProfile string `json:"id_profile" db:"id_profile"`
}
