package model

type Command struct {
	Id       int    `gorm:"PRIMARY_KEY"`
	Word     string `gorm:"unique;not null"`
	Response string `gorm:"not null"`
	Prefix   bool   `gorm:"default:'true'"`
}
