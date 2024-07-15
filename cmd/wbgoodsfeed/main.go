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
	"github.com/urfave/cli/v2"
)

func exec(configPath string, withCron bool) error {
	config, err := readConfig(configPath)
	if err != nil {
		return fmt.Errorf("can't read config file")
	}

	fetcher := wildberries.NewFetcher()

	chooser := &product.DefaultChooser{}

	poster, err := telegram.NewPoster(config.Poster.Options.BotTokenEnv, config.Poster.Options.ChannelName)
	if err != nil {
		log.Panic(fmt.Errorf("can't create poster: %w", err))
	}

	pp := usecases.NewPostProduct(fetcher, chooser, poster)

	if !withCron {
		log.Printf("Find product\n")
		err = pp.Do()
		if err != nil {
			return fmt.Errorf("can't post product: %w", err)
		}
		log.Printf("Product posted\n")
	} else {
		c := cron.New()
		_, err = c.AddFunc(config.Interval, func() {
			log.Printf("Find product\n")
			err = pp.Do()
			if err != nil {
				log.Printf("Can't post product: %s\n", err)
				return
			}
			log.Printf("Product posted\n")
		})
		if err != nil {
			return fmt.Errorf("can't add cron job: %w", err)
		}
		c.Run()
	}

	return nil
}

func main() {
	execFlags := []cli.Flag{
		&cli.StringFlag{
			Name:     "config",
			Aliases:  []string{"c"},
			Required: true,
		},
	}

	app := &cli.App{
		Name:  "WB Goods Feed",
		Usage: "публикация интересных товаров Wildberries в соцсетях",
		Commands: []*cli.Command{
			{
				Name:  "post",
				Usage: "Опубликовать пост",
				Flags: execFlags,
				Action: func(ctx *cli.Context) error {
					return exec(ctx.String("config"), false)
				},
			},
			{
				Name:  "run",
				Usage: "Запустить публикацию по интервалу",
				Flags: execFlags,
				Action: func(ctx *cli.Context) error {
					return exec(ctx.String("config"), true)
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Panic(err)
	}
}
