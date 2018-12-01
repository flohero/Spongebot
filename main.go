package main

import (
	"github.com/flohero/Spongebot/bot"
	"github.com/flohero/Spongebot/database"
	"time"
)

func main() {
	persistence := database.InitDb()
	go bot.Listen("token", persistence)
	for {
		time.Sleep(time.Minute)
	}

}
