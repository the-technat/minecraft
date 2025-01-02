package handlers

import (
	"context"
	"fmt"
	"log"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	flygo "github.com/superfly/fly-go"
)

func listServerHandler(bot *telego.Bot, message telego.Message, flyClient *flygo.Client) {

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
	}
	_, _ = bot.SendMessage(tu.MessageWithEntities(
		tu.ID(message.Chat.ID),
		tu.Entity("Servers:"),
		tu.Entity("\n"),
		tu.Entity(fmt.Sprintf("%+v", apps[1].Hostname)).Code(),
	))

}
