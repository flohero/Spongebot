package main

import (
	"github.com/flohero/Spongebot/api"
	"github.com/flohero/Spongebot/bot"
	"github.com/flohero/Spongebot/database"
	"os"
)

func main() {
	persistence := database.InitDb()
	go bot.Listen(os.Getenv("token"), persistence)
	api.Serve(persistence)

}
