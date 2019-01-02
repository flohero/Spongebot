package api

import (
	"github.com/flohero/Spongebot/database"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

type Controller struct {
	persistence *database.Persistence
}

func Serve(persistence *database.Persistence) {
	c := &Controller{persistence: persistence}
	r := mux.NewRouter()
	// API
	r.Use(corsAndContentTypeHeader)
	r.Use(c.JwtAuthentication)

	r.HandleFunc("/api/commands", c.GetAllCommands).Methods("GET")
	r.HandleFunc("/api/commands", c.CreateCommand).Methods("POST")
	r.HandleFunc("/api/commands/{id}", c.GetCommandById).Methods("GET")

	r.HandleFunc("/api/configs", c.GetAllConfigs).Methods("GET")
	r.HandleFunc("/api/configs", c.CreateConfig).Methods("POST")
	r.HandleFunc("/api/configs/{id}", c.GetConfigById).Methods("GET")

	r.HandleFunc("/api/user/new", c.CreateAccount).Methods("POST")
	r.HandleFunc("/api/user/login", c.Authenticate).Methods("POST")

	// CORS

	r.HandleFunc("/api/commands", c.options).Methods("OPTIONS")
	r.HandleFunc("/api/commands", c.options).Methods("OPTIONS")
	r.HandleFunc("/api/commands/{id}", c.options).Methods("OPTIONS")

	r.HandleFunc("/api/configs", c.options).Methods("OPTIONS")
	r.HandleFunc("/api/configs", c.options).Methods("OPTIONS")
	r.HandleFunc("/api/configs/{id}", c.options).Methods("OPTIONS")

	r.HandleFunc("/api/user/new", c.options).Methods("OPTIONS")
	r.HandleFunc("/api/user/login", c.options).Methods("OPTIONS")

	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	panic(http.ListenAndServe(":8080", loggedRouter))

}

func (c *Controller) options(w http.ResponseWriter, r *http.Request) {
	w.Header().Del("Content-Type")
	println(w.Header().Get("Content-Type"))
	w.WriteHeader(200)
}
