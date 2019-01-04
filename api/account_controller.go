package api

import (
	"encoding/json"
	"errors"
	"github.com/flohero/Spongebot/database/model"
	"net/http"
)

func (c *Controller) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var temp *model.Account
	err := json.NewDecoder(r.Body).Decode(&temp)
	if err != nil {
		badRequest(w, err)
		return
	}
	err, temp = c.persistence.CreateAccount(temp)
	if err != nil {
		badRequest(w, err)
		return
	}
	created(w)
	writeJson(w, temp)
}

func (c *Controller) Authenticate(w http.ResponseWriter, r *http.Request) {
	account := &model.Account{}
	err := json.NewDecoder(r.Body).Decode(account) //decode the request body into struct and failed if any error occur
	if err != nil {
		badRequest(w, err)
		return
	}

	if account.Username == "" || account.Password == "" {
		badRequest(w, errors.New("No username or password supplied"))
		return
	}

	err, account = c.persistence.Login(account.Username, account.Password)
	if err != nil {
		forbidden(w, err)
		return
	}
	writeJson(w, account)
}

func (c *Controller) GetAllAccounts(w http.ResponseWriter, r *http.Request) {
	accs, err := c.persistence.FindAllAccounts()
	if err != nil {
		internalError(w, err)
		return
	}
	writeJson(w, accs)
}

func (c *Controller) DeleteAccountById(w http.ResponseWriter, r *http.Request) {
	check, id := parseId(w, r)
	if !check {
		return
	}
	c.persistence.DeleteAccountById(id)
	w.WriteHeader(204)
}
