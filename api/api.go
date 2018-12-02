package api

import (
	"encoding/json"
	"errors"
	"github.com/flohero/Spongebot/database"
	"github.com/flohero/Spongebot/database/model"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Server struct {
	persistence *database.Persistence
}

func Serve(persistence *database.Persistence) {
	s := &Server{persistence: persistence}
	r := mux.NewRouter()
	// API
	r.Use(corsAndContentTypeHeader)
	r.HandleFunc("/api/commands", s.GetAllCommands).Methods("GET")
	r.HandleFunc("/api/commands", s.CreateCommand).Methods("POST")
	r.HandleFunc("/api/commands/{id}", s.GetCommandById).Methods("GET")
	panic(http.ListenAndServe(":8080", r))
}

func (s *Server) GetAllCommands(writer http.ResponseWriter, request *http.Request) {
	cmds, err := s.persistence.FindAllCommands()
	if err != nil {
		internalError(writer, err)
		return
	}
	writeJson(writer, cmds)
}

func (s *Server) GetCommandById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		badRequest(w, err)
		return
	}
	cmd := s.persistence.FindCommandById(id)
	if cmd.Id == 0 {
		notFound(w, errors.New("not found"))
		return
	}
	writeJson(w, cmd)
}

func (s *Server) CreateCommand(w http.ResponseWriter, r *http.Request) {
	var cmd model.Command
	err := json.NewDecoder(r.Body).Decode(&cmd)
	if err != nil {
		badRequest(w, err)
		return
	}
	c := s.persistence.FindCommandByWord(cmd.Word)
	if c == nil {
		c := &model.Command{Word: cmd.Word, Response: cmd.Response, Prefix: cmd.Prefix}
		s.persistence.CreateCommand(c)
	}
	writeJson(w, &c)
}

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

func writeError(w http.ResponseWriter, err error, code int) {
	w.WriteHeader(code)
	w.Write([]byte(err.Error()))
}

func corsAndContentTypeHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}
