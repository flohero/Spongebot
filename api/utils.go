package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/flohero/Spongebot/database/model"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"strings"
)

func writeJson(w http.ResponseWriter, obj interface{}) {
	if err := json.NewEncoder(w).Encode(obj); err != nil {
		internalError(w, err)
	}
}

func internalError(w http.ResponseWriter, err error) {
	writeError(w, err, 500)
}

func badRequest(w http.ResponseWriter, err error) {
	writeError(w, err, 400)
}

func notFound(w http.ResponseWriter, err error) {
	writeError(w, err, 404)
}

func forbidden(w http.ResponseWriter, err error) {
	writeError(w, err, 403)
}

func writeError(w http.ResponseWriter, err error, code int) {
	w.WriteHeader(code)
	w.Write([]byte(err.Error()))
}

func corsAndContentTypeHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		if r.Method == "OPTIONS" {
			w.Header().Add("Access-Control-Allow-Headers", "Content-Type, authorization")
			w.WriteHeader(200)
			return
		}
		if strings.HasPrefix(r.RequestURI, "/static/") {
			w.Header().Add("Content-Type", "text/html; charset=utf-8")
			return
		}
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func getIdFromPath(w http.ResponseWriter, r *http.Request) (int, error) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return -1, err
	}
	return id, nil
}

func created(w http.ResponseWriter) {
	w.WriteHeader(201)
}

func isRegexOk(regex string) bool {
	if regex == "*" || regex == "°*" || regex == "+" {
		return false
	}
	return true
}

func regexNotOk(w http.ResponseWriter) {
	badRequest(w, errors.New(fmt.Sprint("Regular expressions like '*', '^*' or '+' are not allowed.")))
}

func commandValid(cmd model.Command, w http.ResponseWriter) bool {
	if cmd.Regex == "" || cmd.Response == "" || cmd.Description == "" {
		malformedBody(w)
		return false
	}
	if !isRegexOk(cmd.Regex) {
		regexNotOk(w)
		return false
	}
	return true
}

func malformedId(w http.ResponseWriter) {
	badRequest(w, errors.New(fmt.Sprint("Malformed ID or not provided")))
}

func parseCommand(w http.ResponseWriter, r *http.Request) (bool, *model.Command) {
	var temp model.Command
	err := json.NewDecoder(r.Body).Decode(&temp)
	if err != nil {
		malformedBody(w)
		return false, nil
	}
	if !commandValid(temp, w) {
		return false, nil
	}
	return true, &temp
}

func malformedBody(w http.ResponseWriter) {
	badRequest(w, errors.New("Malformed body"))
}
