package model

type Command struct {
	Id          int    `gorm:"PRIMARY_KEY"json:"id"`
	Regex       string `gorm:"unique;not null"json:"regex"`
	Description string `gorm:""json:"description"`
	Response    string `gorm:"not null"json:"response"`
	Script      bool   `gorm:"not null"json:"script"`
}
