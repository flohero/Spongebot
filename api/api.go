package api

import (
	"encoding/json"
	"github.com/flohero/Spongebot/database"
	"github.com/gorilla/mux"
	"net/http"
)
type Server struct {
	persistence *database.Persistence
}
func Serve(persistence *database.Persistence) {
	s := &Server{persistence: persistence}
	r := mux.NewRouter()
	// API
	r.Use(commonMiddleware)
	r.HandleFunc("/api/commands", s.GetAllCommands).Methods("GET")
	panic(http.ListenAndServe(":8080", r))
}

func (s *Server) GetAllCommands(writer http.ResponseWriter, request *http.Request) {
	cmds := s.persistence.FindCommandByWord("ping")
	if json.NewEncoder(writer).Encode(cmds) != nil {
		internalError(writer)
	}
}

func internalError(w http.ResponseWriter) {
	w.WriteHeader(500)
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
