package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/flohero/Spongebot/database/model"
	"github.com/starlight-go/starlight"
)

func (b *Bot) execScript(message string, cmd *model.Command, session *discordgo.Session, msg *discordgo.MessageCreate) (string, error) {
	s := &scriptResult{
		Message:  message,
		GuildId:  msg.GuildID,
		AuthorId: msg.Author.ID,
	}
	globals := map[string]interface{}{
		"s":        s,
		"kickUser": session.GuildMemberDeleteWithReason,
	}
	_, err := starlight.Eval([]byte(cmd.Response), globals, nil)
	if err != nil {
		return "", err
	}
	return s.Result, nil
}

func (b *Bot) respondToMessage(session *discordgo.Session, msg *discordgo.MessageCreate, content string) {
	sendingMessageError(session.ChannelMessageSend(msg.ChannelID, content))
}

func sendingMessageError(msg *discordgo.Message, err error) {
	if err != nil {
		fmt.Printf("Error sending message: %s\n%s", msg.Content, err)
	}
}
