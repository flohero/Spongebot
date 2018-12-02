package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/flohero/Spongebot/database"
	"strings"
)

const prefix string = "_"

type Bot struct {
	persistence *database.Persistence
}

func Listen(token string, persistence *database.Persistence) {
	discord, err := discordgo.New(fmt.Sprintf("Bot %s", token))
	if err != nil {
		panic(err)
	}
	defer discord.Close()
	if err = discord.Open(); err != nil {
		panic(err)
	}
	bot := Bot{persistence: persistence}
	discord.AddHandler(bot.onMessage)
}

func (b *Bot) onMessage(session *discordgo.Session, msg *discordgo.MessageCreate) {
	if msg.Author.ID == session.State.User.ID {
		return
	}
	cmdStr := strings.Split(msg.Content, " ")[0]
	if cmd := b.persistence.FindCommandByWord(cmdStr); cmd.Id != 0 {
		b.respondToMessage(session, msg, cmd.Response)
	}
}

func (b *Bot) respondToMessage(session *discordgo.Session, msg *discordgo.MessageCreate, content string) {
	_, err := session.ChannelMessageSend(msg.ChannelID, content)
	if err != nil {
		fmt.Printf("Error sending message: %s", err)
	}
}
