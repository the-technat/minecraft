package handlers

import (
	"context"
	"fmt"
	"log"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	flygo "github.com/superfly/fly-go"
	"github.com/superfly/fly-go/flaps"
	"github.com/the-technat/telegram-minecraft-manager/pkg/config"
)

func stopHandler(bot *telego.Bot, message telego.Message, flyClient *flygo.Client, C *config.Config) {

	_, _, args := tu.ParseCommand(message.Text)

	// check if we have a server name
	if len(args) != 0 {
		server := args[0]
		if server != "" {
			srvApp, err := flyClient.GetApp(context.TODO(), fmt.Sprintf("%s%s", C.FlyAppPrefix, server))
			if err == nil {
				_, _ = bot.SendMessage(tu.MessageWithEntities(
					tu.ID(message.Chat.ID),
					tu.Entity("Trying to stop "),
					tu.Entity(server),
					tu.Entity("..."),
				))

				ips, _ := flyClient.GetIPAddresses(context.TODO(), srvApp.Name)
				for _, ip := range ips {
					flyClient.ReleaseIPAddress(context.TODO(), srvApp.Name, ip.Address)
					_, _ = bot.SendMessage(tu.MessageWithEntities(
						tu.ID(message.Chat.ID),
						tu.Entity("Removing IP "),
						tu.Entity(ip.Address),
						tu.Entity(" from server "),
						tu.Entity(server),
					))
				}

				// org, _ := flyClient.GetOrganizationByApp(context.TODO(), srvApp.Name)
				// token, _ := flyClient.CreateDelegatedWireGuardToken(context.TODO(), org, srvApp.Name)
				// initialize fly flaps client
				flapsClient, err := flaps.NewWithOptions(context.TODO(), flaps.NewClientOpts{
					AppName: srvApp.Name,
				})
				log.Printf(err.Error())

				machines, _ := flapsClient.ListActive(context.TODO())
				log.Printf("%+v", machines)
				err = flapsClient.Stop(context.TODO(), flygo.StopMachineInput{ID: machines[0].ID}, "")
				_, _ = bot.SendMessage(tu.MessageWithEntities(
					tu.ID(message.Chat.ID),
					tu.Entity("Stopping server failed: "),
					tu.Entity(err.Error()).Code(),
				))

				return

			}
			_, _ = bot.SendMessage(tu.MessageWithEntities(
				tu.ID(message.Chat.ID),
				tu.Entity("That server doesn't seem to exist:"),
				tu.Entity("\n"),
				tu.Entity(err.Error()).Code(),
			))
			return

		}
	}

	_, _ = bot.SendMessage(tu.MessageWithEntities(
		tu.ID(message.Chat.ID),
		tu.Entity("/stop needs the name of a server to stop. "),
		tu.Entity("Try again with a name. "),
		tu.Entity("You can find all servers by using "),
		tu.Entity("/listServers").BotCommand(),
	))
	return

	// if err != nil {
	// 	log.Printf("Error retrieving app list: %q", err)
	// 	_, _ = bot.SendMessage(tu.MessageWithEntities(
	// 		tu.ID(message.Chat.ID),
	// 		tu.Entity("Error listing servers:"),
	// 		tu.Entity("\n"),
	// 		tu.Entity(err.Error()).Code(),
	// 	))
	// 	return
	// }

	// servers := []string{}
	// for _, app := range apps {
	// 	if strings.HasPrefix(app.Name, C.FlyAppPrefix) {
	// 		name := app.Name
	// 		name = strings.TrimPrefix(name, C.FlyAppPrefix)
	// 		servers = append(servers, name)
	// 	}
	// }

	// _, _ = bot.SendMessage(tu.MessageWithEntities(
	// 	tu.ID(message.Chat.ID),
	// 	tu.Entity("Servers: "),
	// 	tu.Entity(fmt.Sprintf("%s", servers)).Italic(),
	// ))

}
