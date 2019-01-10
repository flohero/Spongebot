package bot

import (
	"errors"
	"github.com/bwmarrin/discordgo"
	"time"
)

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

func (b *Bot) sendEmbed(session *discordgo.Session, msg *discordgo.MessageCreate, embed *discordgo.MessageEmbed) {
	sendingMessageError(session.ChannelMessageSendEmbed(msg.ChannelID, embed))
}

func (b *Bot) sendError(session *discordgo.Session, msg *discordgo.MessageCreate, err error) {
	embed := buildEmbed(0xFF0000, "Error", err.Error(), nil, nil, nil)
	b.sendEmbed(session, msg, embed)
}
