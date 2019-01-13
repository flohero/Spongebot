package api

import (
	"github.com/flohero/Spongebot/channel"
	"net/http"
)

func (c *Controller) StopDiscordBot(w http.ResponseWriter, r *http.Request) {
	c.stopBotChan <- channel.StopFlag{}
}
