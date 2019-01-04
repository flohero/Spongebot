package model

import "github.com/dgrijalva/jwt-go"

type Token struct {
	UserId int
	Admin  bool
	Claims jwt.StandardClaims
}

func (t *Token) Valid() error {
	return t.Claims.Valid()
}

type Account struct {
	Id       int    `gorm:"PRIMARY_KEY"json:"id,omitempty"`
	Username string `gorm:"unique;not null"json:"username,omitempty"`
	Password string `gorm:"not null"json:"password,omitempty"`
	Admin    bool   `gorm:"not null"json:"admin,omitempty"`
	Token    string `sql:"-"json:"token,omitempty"`
}
