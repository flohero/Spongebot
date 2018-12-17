package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/flohero/Spongebot/database"
	"github.com/flohero/Spongebot/database/model"
	"strings"
)

const prefix string = "_"

type Bot struct {
	persistence *database.Persistence
}

func Listen(config *model.Config, persistence *database.Persistence) {
	discord, err := discordgo.New(fmt.Sprintf("Bot %s", config.Token))
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
	if len([]rune(cmdStr)) <= 0 {
		return
	}
	prf := string([]rune(cmdStr)[0]) == prefix
	command := cmdStr
	if prf {
		command = cmdStr[len(cmdStr)-(len(cmdStr)-1):]
	}
	if cmd := b.persistence.FindCommandByWord(command); cmd.Id != 0 && cmd.Prefix == prf {
		b.respondToMessage(session, msg, cmd.Response)
	}
}

func (b *Bot) respondToMessage(session *discordgo.Session, msg *discordgo.MessageCreate, content string) {
	_, err := session.ChannelMessageSend(msg.ChannelID, content)
	if err != nil {
		fmt.Printf("Error sending message: %s", err)
	}
}
