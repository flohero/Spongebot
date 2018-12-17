package model

type Config struct {
	Id     int    `gorm:"PRIMARY_KEY"json:"id"`
	Token  string `gorm:"unique;not null"json:"word"`
	Prefix string `json:"prefix"`
}
