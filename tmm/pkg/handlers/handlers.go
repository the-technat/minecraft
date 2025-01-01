package handlers

import (
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
	"github.com/the-technat/telegram-minecraft-manager/pkg/config"
)

// RegisterHandlers tells the bot what do to in which case
func RegisterHandlers(bh *th.BotHandler, C *config.Config) {
	// register middleware that will determine if a user is allowed to execute that command
	bh.Use(
		func(bot *telego.Bot, update telego.Update, next th.Handler) {
			if AuthMiddleware(bot, update, C) {
				next(bot, update)
			} else {
				UnauthorizedHandler(bot, update)
			}
		},
	)

	// bh.Handle(stopHandler, th.CommandEqual("/stop"))

	// Register new handler with match on command `/help`
	bh.HandleMessage(func(bot *telego.Bot, message telego.Message) {
		HelpHandler(bot, message)
	}, th.CommandEqual("help"))

	// Handle all other messages
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		_, _ = bot.SendMessage(tu.Message(update.Message.Chat.ChatID(), "I'm managing Minecraft Servers, to learn how to use me, type /help. If you want more information about my creator, see https://github.com/the-technat/telegram-minecraft-manager"))
	}, th.AnyMessage())

}

func UnauthorizedHandler(bot *telego.Bot, update telego.Update) {
	_, _ = bot.SendMessage(tu.Message(update.Message.Chat.ChatID(), "You are not authorized for this command."))
}

// func messageHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
// 	if C.IsAdmin(update.Message.From.Username) {
// 		conn := new(mcrcon.MCConn)
// 		err := conn.Open(fmt.Sprintf("%s:%d", C.Servers[0].RCONHost, C.Servers[0].RCONPort), C.Servers[0].RCONPassword)
// 		if err != nil {
// 			log.Fatalln("Open failed", err)
// 		}
// 		defer conn.Close()

// 		err = conn.Authenticate()
// 		if err != nil {
// 			log.Fatalln("Auth failed", err)
// 		}

// 		resp, err := conn.SendCommand("tps")
// 		if err != nil {
// 			log.Fatalln("Command failed", err)
// 		}
// 		b.SendMessage(ctx, &bot.SendMessageParams{
// 			ChatID:    update.Message.Chat.ID,
// 			Text:      resp,
// 			ParseMode: models.ParseModeMarkdown,
// 		})
// 		log.Println(resp)
// 	}
// }
