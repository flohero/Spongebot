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
	p.db.DropTable(&model.Command{})
	p.db.DropTable(&model.Account{})
	p.db.CreateTable(&model.Command{})
	p.db.CreateTable(&model.Config{})
	p.db.CreateTable(&model.Account{})
}

func (p *Persistence) initData() {
	p.CreateCommand(&model.Command{Word: "ping", Response: "pong", Script: false, Prefix: false})
	p.CreateCommand(&model.Command{Word: "peng", Response: "s.Result = s.Message.upper()", Script: true, Prefix: false})
	s := "s.Result = \"|\".join(s.Message.split(\" \"))"
	p.CreateCommand(&model.Command{Word: "hello", Response: s, Script: true, Prefix: false})
	p.CreateAccount(&model.Account{Username: "sponge", Password: "bot"})
}
