package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	"time"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
	flygo "github.com/superfly/fly-go"

	"github.com/the-technat/telegram-minecraft-manager/pkg/config"
	"github.com/the-technat/telegram-minecraft-manager/pkg/handlers"
)

func main() {
	// Load config according to schema
	C := config.LoadConfig()

	// create new bot instance
	bot, err := telego.NewBot(C.Token, telego.WithDefaultLogger(C.Debug, true))
	if err != nil {
		log.Fatal(err.Error())
	}
	// check if token is valid
	botUser, err := bot.GetMe()
	if err != nil {
		log.Fatalf("Create bot: %q", err)
	}
	log.Printf("Running as user: %s", botUser.Username)

	// Get updates & update webhook url to send updates to
	updates, err := bot.UpdatesViaWebhook("/", telego.WithWebhookSet(tu.Webhook(C.WebhookURL).WithSecretToken(C.WebhookToken)))
	if err != nil {
		log.Fatalf("Updates via webhook: %q", err)
	}

	// Create bot handler with stop timeout
	bh, err := th.NewBotHandler(bot, updates)
	if err != nil {
		log.Fatalf("Bot handler: %q", err)
	}

	// initalize flyClient
	flyClient := flygo.NewClientFromOptions(flygo.ClientOptions{
		AccessToken:      C.FlyOrgToken,
		EnableDebugTrace: &C.Debug,
		BaseURL:          "https://api.fly.io",
	})
	if !flyClient.Authenticated() {
		log.Fatal("Fly API Token is invalid")
	}

	// register handlers to act on updates
	handlers.RegisterHandlers(bh, C, flyClient)

	// Initialize signal handling
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)

	// Initialize done chan
	done := make(chan struct{}, 1)

	// Handle stop signal (Ctrl+C)
	go func() {
		// Wait for stop signal
		<-sigs

		fmt.Println("Stopping...")

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		err = bot.StopWebhookWithContext(ctx)
		if err != nil {
			log.Printf("Failed to stop webhook properly: %q", err)
		}

		fmt.Println("Webhook done")

		bh.StopWithContext(ctx)
		log.Print("Bot handler done")

		// Notify that stop is done
		done <- struct{}{}
	}()

	// Start handling in goroutine
	go bh.Start()
	log.Print("Ready to handle messages")

	// Start server for receiving requests from the Telegram
	go func() {
		err = bot.StartWebhook(fmt.Sprintf("0.0.0.0:%d", C.Port))
		if err != nil {
			log.Fatalf("Failed to start webhook: %q", err)
		}
	}()

	// Wait for the stop process to be completed
	<-done
	log.Print("Done")
}
