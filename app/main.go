package main

import (
	"fmt"
	tbapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkarpovich/tg-link-keeper-bot/app/bot/linkstore"
	"github.com/pkarpovich/tg-link-keeper-bot/app/config"
	"github.com/pkarpovich/tg-link-keeper-bot/app/events"
	"log"
)

func main() {
	log.Printf("[INFO] starting app...")

	cfg, err := config.Init()
	if err != nil {
		log.Fatalf("[ERROR] %v", err)
	}

	if err := execute(cfg); err != nil {
		log.Printf("[ERROR] %v", err)
	}
}

func execute(config *config.Config) error {
	tbAPI, err := tbapi.NewBotAPI(config.Telegram.Token)
	if err != nil {
		return fmt.Errorf("failed to create Telegram events: %w", err)
	}

	linkdingClient := linkstore.NewLinkStoreClient(config.LinkStore.Url)

	tgListener := &events.TelegramListener{
		SuperUsers: config.Telegram.SuperUsers,
		TbAPI:      tbAPI,
		Bot:        linkdingClient,
	}

	if err := tgListener.Do(); err != nil {
		return fmt.Errorf("failed to start Telegram listener: %w", err)
	}

	return nil
}
