package main

import (
	"github.com/flohero/Spongebot/api"
	"github.com/flohero/Spongebot/bot"
	"github.com/flohero/Spongebot/channel"
	"github.com/flohero/Spongebot/database"
	"github.com/flohero/Spongebot/database/model"
	"github.com/flohero/Spongebot/utils"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}
	utils.InitEnvironment()
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
		token = utils.Environment["DISCORD_TOKEN"]
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
