package database

import (
	"fmt"
	"github.com/flohero/Spongebot/database/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
)

type Persistence struct {
	db *gorm.DB
}

func InitDb() *Persistence {
	db, err := connectToPostgres()
	if err != nil {
		panic(fmt.Sprintf("Error opening DB: %s", err.Error()))
	}
	p := &Persistence{db: db}
	p.createDB()
	p.initData()
	return p
}

func connectToPostgres() (*gorm.DB, error) {
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	pw := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	return gorm.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbName, pw))
}

func connectToSqlite() (*gorm.DB, error) {
	return gorm.Open("sqlite3", "bot.db")
}

func (p *Persistence) createDB() {
	//p.db.DropTable(&model.Command{})
	p.db.DropTable(&model.Account{})
	p.db.CreateTable(&model.Command{})
	p.db.CreateTable(&model.Config{})
	p.db.CreateTable(&model.Account{})
}

func (p *Persistence) initData() {
	p.CreateCommand(&model.Command{Regex: "^ping", Description: "Will response with pong.", Response: "pong", Script: false})
	p.CreateCommand(&model.Command{Regex: "peng", Description: "This will make your message uppercase.", Response: "s.Result = s.Message.upper()", Script: true})
	s := "s.Result = \"|\".join(s.Message.split(\" \"))"
	p.CreateCommand(&model.Command{Regex: ".+\\s", Description: "This will replace all whitespaces with a pipe.", Response: s, Script: true})
	p.CreateAccount(&model.Account{Username: "sponge", Password: "bot"})
}
