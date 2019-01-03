package model

import "github.com/dgrijalva/jwt-go"

type Token struct {
	UserId int
	Claims jwt.StandardClaims
}

func (t *Token) Valid() error {
	return t.Claims.Valid()
}

type Account struct {
	Id       int    `gorm:"PRIMARY_KEY"json:"id"`
	Username string `gorm:"unique;not null"json:"username"`
	Password string `gorm:"not null"json:"password"`
	Token    string `sql:"-"json:"token"`
}
