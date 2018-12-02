package api

import (
	"encoding/json"
	"github.com/flohero/Spongebot/database"
	"github.com/gorilla/mux"
	"net/http"
)

type Controller struct {
	persistence *database.Persistence
}

func Serve(persistence *database.Persistence) {
	c := &Controller{persistence: persistence}
	r := mux.NewRouter()
	// API
	r.Use(corsAndContentTypeHeader)
	r.HandleFunc("/api/commands", c.GetAllCommands).Methods("GET")
	r.HandleFunc("/api/commands", c.CreateCommand).Methods("POST")
	r.HandleFunc("/api/commands/{id}", c.GetCommandById).Methods("GET")
	panic(http.ListenAndServe(":8080", r))
}

func writeJson(w http.ResponseWriter, obj interface{}) {
	if err := json.NewEncoder(w).Encode(obj); err != nil {
		internalError(w, err)
	}
}

func loadJson(obj interface{}) {

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
