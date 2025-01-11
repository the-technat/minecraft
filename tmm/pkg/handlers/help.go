package handlers

import (
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func HelpHandler(bot *telego.Bot, message telego.Message) {
	_, _ = bot.SendMessage(tu.Messagef(
		tu.ID(message.Chat.ID),
		"So you want to know how to use me %s?", message.From.FirstName,
	))
	_, _ = bot.SendMessage(tu.Message(
		tu.ID(message.Chat.ID),
		"I'm listening for commands, they always start with /"))
	_, _ = bot.SendMessage(tu.Message(
		tu.ID(message.Chat.ID),
		"Each command will tell you if you're using it wrong, so don't be shy to run one of the supported commands"))
	_, _ = bot.SendMessage(tu.Message(
		tu.ID(message.Chat.ID),
		"I'm currently supporting these commands: /listServers, /stop <server> and /start <server>"))
}
