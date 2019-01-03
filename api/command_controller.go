package api

import (
	"errors"
	"github.com/flohero/Spongebot/database/model"
	"net/http"
)

// GetAllCommands takes a http.ResponseWriter and a *http.Request.
// It will write all Commands in JSON Format to writer
func (c *Controller) GetAllCommands(writer http.ResponseWriter, request *http.Request) {
	cmds, err := c.persistence.FindAllCommands()
	if err != nil {
		internalError(writer, err)
		return
	}
	writeJson(writer, cmds)
}

func (c *Controller) GetCommandById(w http.ResponseWriter, r *http.Request) {
	check, id := parseId(w, r)
	if !check {
		return
	}
	cmd := c.persistence.FindCommandById(id)
	if cmd.Id == 0 {
		notFound(w, errors.New("not found"))
		return
	}
	writeJson(w, cmd)
}

func (c *Controller) CreateCommand(w http.ResponseWriter, r *http.Request) {
	check, temp := parseCommand(w, r)
	if !check {
		return
	}
	var cmd *model.Command
	cmd = c.persistence.FindCommandByWord(temp.Regex)
	if cmd.Id == 0 {
		cmd = &model.Command{Regex: temp.Regex, Response: temp.Response, Script: temp.Script, Description: temp.Description}
		c.persistence.CreateCommand(cmd)
	}
	created(w)
	writeJson(w, &cmd)
}

func (c *Controller) UpdateCommandById(w http.ResponseWriter, r *http.Request) {
	check, temp := parseCommand(w, r)
	if !check {
		return
	}
	check, id := parseId(w, r)
	if !check {
		return
	}
	temp.Id = id
	c.persistence.UpdateCommand(temp)
	w.WriteHeader(204)
}

func (c *Controller) DeleteCommandById(w http.ResponseWriter, r *http.Request) {
	check, id := parseId(w, r)
	if !check {
		return
	}
	c.persistence.DeleteCommandById(id)
	w.WriteHeader(204)
}

func parseId(w http.ResponseWriter, r *http.Request) (bool, int) {
	id, err := getIdFromPath(w, r)
	if err != nil {
		malformedId(w)
		return false, -1
	}
	return true, id
}
