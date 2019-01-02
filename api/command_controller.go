package api

import (
	"encoding/json"
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
	id, err := getIdFromPath(w, r)
	if err != nil {
		badRequest(w, err)
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
	var temp model.Command
	err := json.NewDecoder(r.Body).Decode(&temp)
	if err != nil || temp.Regex == "" || temp.Response == "" {
		badRequest(w, errors.New("Malformed body"))
		return
	}
	var cmd *model.Command
	cmd = c.persistence.FindCommandByWord(temp.Regex)
	if cmd.Id == 0 {
		cmd = &model.Command{Regex: temp.Regex, Response: temp.Response, Prefix: temp.Prefix}
		c.persistence.CreateCommand(cmd)
	}
	created(w)
	writeJson(w, &cmd)
}
