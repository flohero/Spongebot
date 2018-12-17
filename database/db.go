package database

import (
	"fmt"
	"github.com/flohero/Spongebot/database/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Persistence struct {
	db *gorm.DB
}

func InitDb() *Persistence {
	db, err := gorm.Open("sqlite3", "bot.db")
	if err != nil {
		panic(fmt.Sprintf("Error opening DB: %s", err.Error()))
	}
	p := &Persistence{db: db}
	p.createDB()
	p.initData()
	return p
}

func (p *Persistence) createDB() {
	p.db.CreateTable(&model.Command{})
	p.db.CreateTable(&model.Config{})
	p.db.CreateTable(&model.Account{})
}

func (p *Persistence) initData() {
	p.CreateCommand(&model.Command{Word: "ping", Response: "pong", Prefix: false})
}
