package models

type User struct {
	Id           int    `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Login        string `json:"login" db:"login"`
	Email        string `json:"email" db:"email"`
	Password     string `json:"password" db:"password"`
	FullNameUser string `json:"full_name_user" db:"full_name_user"`
	BirthDate    string `json:"birthday_user" db:"birthday_user"`
	Photo        string `json:"photo" db:"photo"`
}

type Credentials struct {
	Password string `json:"password" db:"password"`
}
