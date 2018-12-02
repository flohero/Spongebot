package api

import (
	"encoding/json"
	"errors"
	"github.com/flohero/Spongebot/database/model"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (c *Controller) GetAllCommands(writer http.ResponseWriter, request *http.Request) {
	cmds, err := c.persistence.FindAllCommands()
	if err != nil {
		internalError(writer, err)
		return
	}
	writeJson(writer, cmds)
}

func (c *Controller) GetCommandById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
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
	if err != nil {
		badRequest(w, err)
		return
	}
	var cmd *model.Command
	cmd = c.persistence.FindCommandByWord(temp.Word)
	if cmd.Id == 0 {
		cmd = &model.Command{Word: temp.Word, Response: temp.Response, Prefix: temp.Prefix}
		c.persistence.CreateCommand(cmd)
	}
	writeJson(w, &cmd)
}
