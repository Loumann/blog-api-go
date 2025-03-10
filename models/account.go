package models

import "github.com/dgrijalva/jwt-go"

type User struct {
	Id           int    `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Login        string `json:"login" db:"login"`
	Email        string `json:"email" db:"email"`
	Password     string `json:"password" db:"password"`
	FullNameUser string `json:"full_name" db:"full_name"`
	Photo        string `json:"photo" db:"photo"`
}

type Credentials struct {
	Password string `json:"password" db:"password"`
}

type Claims struct {
	UserId int `json:"id" db:"id"`
	jwt.StandardClaims
}
