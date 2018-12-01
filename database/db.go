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
}

func (p *Persistence) initData() {
	p.db.Create(&model.Command{Word: "ping", Response: "pong", Prefix: false})
}

// TODO doestn return a value
func (p *Persistence) FindAllCommands() ([]*model.Command, error) {
	rows, err := p.db.Table("commands").Rows()
	if err != nil {
		return nil, err
	}
	commands := make([]*model.Command, 0)
	for rows.Next() {
		cmd := &model.Command{}
		if rows.Scan(&cmd.Id, &cmd.Word, &cmd.Response) != nil {
			return nil, err
		}
		commands = append(commands, cmd)
	}
	return commands, nil
}

func (p *Persistence) FindCommandByWord(word string) (cmd *model.Command) {
	cmd = &model.Command{}
	p.db.Where(&model.Command{Word: word}).First(cmd)
	return cmd
}
