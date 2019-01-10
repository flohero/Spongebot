package main

import (
	"github.com/flohero/Spongebot/api"
	"github.com/flohero/Spongebot/bot"
	"github.com/flohero/Spongebot/channel"
	"github.com/flohero/Spongebot/database"
	"github.com/flohero/Spongebot/database/model"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func main() {
	persistence := database.InitDb()
	var conf *model.Config
	conf, err := persistence.FindFirstActiveConfig()
	var token string
	if err == nil {
		println("Used token from DB")
		token = conf.Token
	} else {
		token = os.Getenv("token")
		if token == "" {
			panic("No token provided")
		}
		println("Used token from Env")
		conf.Token = token
		conf.Prefix = "_"
		conf.Active = true
		persistence.CreateConfig(conf)
	}
	stopChan := make(chan channel.StopFlag)
	go bot.Listen(conf, persistence, stopChan)
	api.Serve(persistence, stopChan)
}
