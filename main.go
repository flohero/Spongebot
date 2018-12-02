package main

import (
	"github.com/flohero/Spongebot/api"
	"github.com/flohero/Spongebot/bot"
	"github.com/flohero/Spongebot/database"
	"os"
)

//TODO Move the init of the bot into the api folder, so users can start the bot from the website
func main() {
	persistence := database.InitDb()
	conf := persistence.FindConfigById(1)
	var token string
	if conf != nil {
		println("Used token from DB")
		token = conf.Token
	} else {
		println("Used token from Env")
		token = os.Getenv("token")
	}
	go bot.Listen(token, persistence)
	api.Serve(persistence)

}
