package bot

import (
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/flohero/Spongebot/database"
	"github.com/flohero/Spongebot/database/model"
	"github.com/starlight-go/starlight"
	"time"
)

const prefix string = "_"

type Bot struct {
	persistence *database.Persistence
	config      *model.Config
}

type scriptResult struct {
	Message  string
	Result   string
	GuildId  string
	AuthorId string
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
	bot := Bot{persistence: persistence, config: config}
	discord.AddHandler(bot.onMessage)
}

func (b *Bot) onMessage(session *discordgo.Session, msg *discordgo.MessageCreate) {
	if msg.Author.ID == session.State.User.ID {
		return
	}

	if cmds, err := b.persistence.FindCommandByRegex(msg.Content); len(cmds) != 0 && err == nil {
		fmt.Printf("Got %v match(es) with message '%s' from User %s", len(cmds), msg.Content, msg.Author.Username)
		for _, cmd := range cmds {
			if cmd.Script {
				res, err := b.execScript(msg.Content, cmd, session, msg)
				if err != nil {
					fmt.Printf("\nError running script: %s", err)
					b.sendError(session, msg, errors.New(fmt.Sprintf("Error running script [id: %v]: %s", cmd.Id, err.Error())))
					return
				}
				if res == "" {
					println("No result")
				}
				b.respondToMessage(session, msg, res)
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

func (b *Bot) sendHelp(session *discordgo.Session, msg *discordgo.MessageCreate) {
	fields := make([]*discordgo.MessageEmbedField, 0)
	cmds, err := b.persistence.FindAllCommands()
	if err != nil {
		b.sendError(session, msg, errors.New("error finding commands"))
		return
	}
	for _, v := range cmds {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:  v.Regex,
			Value: v.Description,
		})
	}
	b.sendEmbed(session, msg,
		buildEmbed(0x00ff00, "HELP",
			"These commands are available. Keep in mind that all commands are regular expressions. https://www.regular-expressions.info/",
			fields, nil, nil),
	)

}

func (b *Bot) respondToMessage(session *discordgo.Session, msg *discordgo.MessageCreate, content string) {
	sendingMessageError(session.ChannelMessageSend(msg.ChannelID, content))
}

func (b *Bot) sendEmbed(session *discordgo.Session, msg *discordgo.MessageCreate, embed *discordgo.MessageEmbed) {
	sendingMessageError(session.ChannelMessageSendEmbed(msg.ChannelID, embed))
}

func (b *Bot) sendError(session *discordgo.Session, msg *discordgo.MessageCreate, err error) {
	embed := buildEmbed(0xFF0000, "Error", err.Error(), nil, nil, nil)
	b.sendEmbed(session, msg, embed)
}

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

func buildEmbed(color int, title, description string,
	fields []*discordgo.MessageEmbedField, image *discordgo.MessageEmbedImage,
	thumbnail *discordgo.MessageEmbedThumbnail) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Author:      &discordgo.MessageEmbedAuthor{},
		Color:       color, // 0x00ff00
		Description: description,
		Fields:      fields,
		Image:       image,
		Thumbnail:   thumbnail,
		Timestamp:   time.Now().Format(time.RFC3339), // Discord wants ISO8601; RFC3339 is an extension of ISO8601 and should be completely compatible.
		Title:       title,
	}
}

func sendingMessageError(msg *discordgo.Message, err error) {
	if err != nil {
		fmt.Printf("Error sending message: %s\n%s", msg.Content, err)
	}
}
