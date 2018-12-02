package api

import (
	"encoding/json"
	"errors"
	"github.com/flohero/Spongebot/database/model"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (c *Controller) GetAllConfigs(w http.ResponseWriter, r *http.Request) {
	configs, err := c.persistence.FindAllConfigs()
	if err != nil {
		internalError(w, err)
		return
	}
	writeJson(w, configs)
}

func (c *Controller) GetConfigById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		badRequest(w, err)
		return
	}
	config := c.persistence.FindConfigById(id)
	if config.Id == 0 {
		notFound(w, errors.New("not found"))
		return
	}
	writeJson(w, config)
}

func (c *Controller) CreateConfig(w http.ResponseWriter, r *http.Request) {
	var temp model.Config
	err := json.NewDecoder(r.Body).Decode(&temp)
	if err != nil {
		badRequest(w, err)
		return
	}
	var config *model.Config
	config = c.persistence.FindConfigByToken(temp.Token)
	if config.Id == 0 {
		config = &model.Config{Token: temp.Token}
		c.persistence.CreateConfig(config)
	}
	writeJson(w, &config)
}
