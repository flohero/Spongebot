package model

type Command struct {
	Id       int    `gorm:"PRIMARY_KEY"json:"id"`
	Word     string `gorm:"unique;not null"json:"word"`
	Response string `gorm:"not null"json:"response"`
	Script   bool   `gorm:"not null"json:"script"`
	Prefix   bool
}
