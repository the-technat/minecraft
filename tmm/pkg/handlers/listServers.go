package handlers

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	flygo "github.com/superfly/fly-go"
	"github.com/the-technat/telegram-minecraft-manager/pkg/config"
)

func listServerHandler(bot *telego.Bot, message telego.Message, flyClient *flygo.Client, C *config.Config) {

	// no idea what role should be
	role := ""
	apps, err := flyClient.GetApps(context.TODO(), &role)

	if err != nil {
		log.Printf("Error retrieving app list: %q", err)
		_, _ = bot.SendMessage(tu.MessageWithEntities(
			tu.ID(message.Chat.ID),
			tu.Entity("Error listing servers:"),
			tu.Entity("\n"),
			tu.Entity(err.Error()).Code(),
		))
		return
	}

	servers := []string{}
	for _, app := range apps {
		if strings.HasPrefix(app.Name, C.FlyAppPrefix) {
			name := app.Name
			name = strings.TrimPrefix(name, C.FlyAppPrefix)
			servers = append(servers, name)
		}
	}

	_, _ = bot.SendMessage(tu.MessageWithEntities(
		tu.ID(message.Chat.ID),
		tu.Entity("Servers: "),
		tu.Entity(fmt.Sprintf("%s", servers)).Italic(),
	))

}
