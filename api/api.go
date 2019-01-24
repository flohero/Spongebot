package api

import (
	"fmt"
	"github.com/flohero/Spongebot/channel"
	"github.com/flohero/Spongebot/database"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

type Controller struct {
	persistence *database.Persistence
	stopBotChan chan channel.StopFlag
}

func Serve(persistence *database.Persistence, stopBotChan chan channel.StopFlag) {
	c := &Controller{persistence: persistence}
	r := mux.NewRouter()
	r2 := mux.NewRouter()
	fs := http.FileServer(http.Dir("./static/"))
	r2.PathPrefix("/").Handler(
		http.StripPrefix("/", fs))

	go runWebsite(r2, "8081")

	// API
	r.Use(corsAndContentTypeHeader)
	r.Use(c.JwtAuthentication)

	r.HandleFunc("/api/commands", c.GetAllCommands).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/commands/new", c.CreateCommand).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/commands/{id}", c.GetCommandById).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/commands/{id}/delete", c.DeleteCommandById).Methods("DELETE", "OPTIONS")
	r.HandleFunc("/api/commands/{id}/update", c.UpdateCommandById).Methods("PUT", "OPTIONS")

	r.HandleFunc("/api/configs", c.GetAllConfigs).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/configs", c.CreateConfig).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/configs/{id}", c.GetConfigById).Methods("GET", "OPTIONS")

	// Admin only
	r.HandleFunc("/api/users", c.GetAllAccounts).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/users/{id}/delete", c.DeleteAccountById).Methods("DELETE", "OPTIONS")
	r.HandleFunc("/api/users/username/{username}", c.GetAccountByUsername).Methods("GET", "OPTIONS")

	r.HandleFunc("/api/user/new", c.CreateAccount).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/user/login", c.Authenticate).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/user/update", c.UpdateAccountById).Methods("PUT", "OPTIONS")

	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	panic(http.ListenAndServe(":8080", loggedRouter))
}

func runWebsite(r *mux.Router, port string) {
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), r)
	if err != nil {
		panic(err)
	}
}
