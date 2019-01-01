package main

import (
	"github.com/flohero/Spongebot/api"
	"github.com/flohero/Spongebot/bot"
	"github.com/flohero/Spongebot/database"
	"github.com/flohero/Spongebot/database/model"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

//TODO Move the init of the bot into the api folder, so users can start the bot from the website
func main() {
	persistence := database.InitDb()
	var conf *model.Config
	conf = persistence.FindConfigById(1)
	var token string
	if conf.Id != 0 {
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
		persistence.CreateConfig(conf)
	}
	go bot.Listen(conf, persistence)
	api.Serve(persistence)
}
