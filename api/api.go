package api

import (
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
	r.Use(c.JwtAuthentication)

	r.HandleFunc("/api/commands", c.GetAllCommands).Methods("GET")
	r.HandleFunc("/api/commands", c.CreateCommand).Methods("POST")
	r.HandleFunc("/api/commands/{id}", c.GetCommandById).Methods("GET")

	r.HandleFunc("/api/configs", c.GetAllConfigs).Methods("GET")
	r.HandleFunc("/api/configs", c.CreateConfig).Methods("POST")
	r.HandleFunc("/api/configs/{id}", c.GetConfigById).Methods("GET")
	panic(http.ListenAndServe(":8080", r))
}
