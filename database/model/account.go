package model

import "github.com/dgrijalva/jwt-go"

type Token struct {
	UserId int
	jwt.StandardClaims
}

type Account struct {
	Id       int    `gorm:"PRIMARY_KEY"json:"id"`
	Email    string `gorm:"unique;not null"json:"email"`
	Password string `gorm:"not null"json:"password"`
	Token    string `sql:"-"json:"token"`
}
