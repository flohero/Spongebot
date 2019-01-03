package database

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/flohero/Spongebot/database/model"
	"log"
)

func (p *Persistence) CreateCommand(cmd *model.Command) {
	p.db.Create(&cmd)
}

func (p *Persistence) FindAllCommands() ([]*model.Command, error) {
	rows, err := p.db.Table("commands").Rows()
	if err != nil {
		return nil, err
	}
	return assignRowsToCommand(rows)
}

func (p *Persistence) FindCommandByWord(word string) (cmd *model.Command) {
	cmd = &model.Command{}
	p.db.Where(&model.Command{Regex: word}).First(cmd)
	return cmd
}

func (p *Persistence) FindCommandById(id int) (cmd *model.Command) {
	cmd = &model.Command{}
	p.db.Where(&model.Command{Id: id}).First(cmd)
	return cmd
}

// Querys the db with a message and if a word or 'regex' matches returns it
func (p *Persistence) FindCommandByRegex(message string) ([]*model.Command, error) {
	rows, err := p.db.Table("commands").Select("*").Where(fmt.Sprintf("'%s' ~ regex", message)).Rows()
	if err != nil {
		return nil, err
	}

	return assignRowsToCommand(rows)
}

func (p *Persistence) DeleteCommandById(id int) {
	p.db.Where(model.Command{Id: id}).Delete(model.Command{})
}

func (p *Persistence) UpdateCommand(cmd *model.Command) {
	p.db.Save(cmd)
}

func assignRowsToCommand(rows *sql.Rows) ([]*model.Command, error) {
	commands := make([]*model.Command, 0)
	for rows.Next() {
		cmd := &model.Command{}
		if err := rows.Scan(&cmd.Id, &cmd.Regex, &cmd.Description, &cmd.Response, &cmd.Script); err != nil {
			log.Fatal(err.Error())
			return nil, errors.New("Error with DB")
		}
		commands = append(commands, cmd)
	}
	return commands, nil
}
