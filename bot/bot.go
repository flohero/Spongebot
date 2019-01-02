package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/flohero/Spongebot/database"
	"github.com/flohero/Spongebot/database/model"
	"github.com/starlight-go/starlight"
)

const prefix string = "_"

type Bot struct {
	persistence *database.Persistence
}

type scriptResult struct {
	Message string
	Result  string
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
	/*cmdStr := strings.Split(msg.Content, " ")[0]
	if len([]rune(cmdStr)) <= 0 {
		return
	}
	prf := string([]rune(cmdStr)[0]) == prefix
	command := cmdStr
	if prf {
		command = cmdStr[len(cmdStr)-(len(cmdStr)-1):]
	}*/
	if cmds, err := b.persistence.FindCommandByRegex(msg.Content); len(cmds) != 0 && err == nil {
		for _, cmd := range cmds {
			if cmd.Script {
				res, err := execScript(msg.Content, cmd)
				if err != nil {
					fmt.Printf("\nError running script: %s", err)
					return
				}
				b.respondToMessage(session, msg, res)
			} else {
				b.respondToMessage(session, msg, cmd.Response)
			}
		}
	} else if err != nil {
		fmt.Println(err)
	}
}

func (b *Bot) respondToMessage(session *discordgo.Session, msg *discordgo.MessageCreate, content string) {
	_, err := session.ChannelMessageSend(msg.ChannelID, content)
	if err != nil {
		fmt.Printf("Error sending message: %s", err)
	}
}

func execScript(message string, cmd *model.Command) (string, error) {
	s := &scriptResult{
		Message: message,
	}
	globals := map[string]interface{}{
		"s": s,
	}
	_, err := starlight.Eval([]byte(cmd.Response), globals, nil)
	if err != nil {
		return "", err
	}
	return s.Result, nil
}
