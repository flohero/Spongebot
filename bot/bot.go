package bot

import (
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/flohero/Spongebot/channel"
	"github.com/flohero/Spongebot/database"
	"github.com/flohero/Spongebot/database/model"
)

var prefix string = "_"

type Bot struct {
	persistence *database.Persistence
	config      *model.Config
	session     *discordgo.Session
}

type scriptResult struct {
	Message  string
	Result   string
	GuildId  string
	AuthorId string
}

func Listen(config *model.Config, persistence *database.Persistence, stopChan chan channel.StopFlag) {
	prefix = config.Prefix
	discord, err := discordgo.New(fmt.Sprintf("Bot %s", config.Token))
	if err != nil {
		panic(err)
	}
	defer discord.Close()
	if err = discord.Open(); err != nil {
		panic(err)
	}
	bot := Bot{persistence: persistence, config: config, session: discord}
	discord.AddHandler(bot.onReady)
	discord.AddHandler(bot.onMessage)
	for {
		select {
		case <-stopChan:
			println("Stopping bot...")
			return
		}
	}
}

func (b *Bot) onMessage(session *discordgo.Session, msg *discordgo.MessageCreate) {
	if msg.Author.ID == session.State.User.ID {
		return
	}
	if cmds, err := b.persistence.FindCommandByRegex(msg.Content); len(cmds) != 0 && err == nil {
		fmt.Printf("Got %v match(es) with message '%s' from User %s\n", len(cmds), msg.Content, msg.Author.Username)
		for _, cmd := range cmds {
			if cmd.Script {
				res, err := b.execScript(msg.Content, cmd, session, msg)
				if err != nil {
					fmt.Printf("\nError running script: %s", err)
					b.sendError(session, msg, errors.New(fmt.Sprintf("Error running script [id: %v]: %s", cmd.Id, err.Error())))
					return
				} else if res == "" {
					println("No result")
				} else {
					b.respondToMessage(session, msg, res)
				}
			} else {
				b.respondToMessage(session, msg, cmd.Response)
			}
		}
	} else if err != nil {
		b.sendError(session, msg, err)
	}
	if msg.Content == fmt.Sprintf("%shelp", prefix) {
		b.sendHelp(session, msg)
	}
}

func (b *Bot) onReady(discord *discordgo.Session, ready *discordgo.Ready) {
	err := discord.UpdateStatus(0, "_help for help")
	if err != nil {
		println("Error setting status")
	}
	fmt.Printf("Bot started on %v servers\n", len(discord.State.Guilds))
}
