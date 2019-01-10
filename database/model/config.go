package model

type Config struct {
	Id     int    `gorm:"PRIMARY_KEY"json:"id"`
	Token  string `gorm:"unique;not null"json:"word"`
	Active bool   `gorm:"unique"json:"active"`
	Prefix string `json:"prefix"`
}
