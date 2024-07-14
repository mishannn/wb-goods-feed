package main

import (
	"fmt"
	"log"
	"os"

	"example.org/wbsniper/internal/entities/product"
	"example.org/wbsniper/internal/integrations/telegram"
	"example.org/wbsniper/internal/integrations/wildberries"
	"example.org/wbsniper/internal/usecases"

	"github.com/robfig/cron/v3"
)

func main() {
	// Classic bot token
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		log.Panic("TELEGRAM_BOT_TOKEN env var is not set\n")
	}

	// Format: @channel_name
	channelName := os.Getenv("TELEGRAM_CHANNEL_NAME")
	if botToken == "" {
		log.Panic("TELEGRAM_CHANNEL_NAME env var is not set\n")
	}

	// CRON format, e.g. 0 * * * *
	postPublishPeriod := os.Getenv("POST_PUBLISH_PERIOD")
	if botToken == "" {
		log.Panic("POST_PUBLISH_PERIOD env var is not set\n")
	}

	fetcher := wildberries.NewFetcher()

	chooser := &product.DefaultChooser{}

	poster, err := telegram.NewPoster(botToken, channelName)
	if err != nil {
		log.Panic(fmt.Errorf("can't create poster: %w", err))
	}

	pp := usecases.NewPostProduct(fetcher, chooser, poster)

	c := cron.New()
	c.AddFunc(postPublishPeriod, func() {
		log.Printf("Find product\n")

		err = pp.Do()
		if err != nil {
			log.Printf("Can't post product: %s\n", err)
			return
		}

		log.Printf("Product posted\n")
	})
	c.Run()
}
