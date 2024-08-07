package telegram

import (
	"fmt"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/mishannn/wb-goods-feed/internal/entities/feed"
)

type Poster struct {
	bot    *tgbotapi.BotAPI
	chatID int64
}

func NewPoster(botTokenEnv string, chatID int64) (*Poster, error) {
	// Classic bot token
	botToken := os.Getenv(botTokenEnv)
	if botToken == "" {
		return nil, fmt.Errorf("bot token env var %s is not set", botTokenEnv)
	}

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		return nil, fmt.Errorf("can't initialize telegram bot: %w", err)
	}

	return &Poster{
		bot:    bot,
		chatID: chatID,
	}, nil
}

func (p *Poster) PublishPost(post feed.Post) error {
	text := fmt.Sprintf("*%s*\n\n%s\n\n[Открыть карточку товара](%s)", post.Title, post.Content, post.Link)

	var err error
	if len(post.Images) != 0 {
		var photosLimit int
		if len(post.Images) >= 10 {
			photosLimit = 10
		} else {
			photosLimit = len(post.Images)
		}

		photos := make([]interface{}, 0, photosLimit)
		for imageIndex, image := range post.Images[0:photosLimit] {
			photo := tgbotapi.NewInputMediaPhoto(tgbotapi.FileURL(image.URL))
			if imageIndex == 0 {
				photo.Caption = text
				photo.ParseMode = "Markdown"
			}
			photos = append(photos, photo)
		}
		mediaGroup := tgbotapi.NewMediaGroup(p.chatID, photos)
		_, err = p.bot.SendMediaGroup(mediaGroup)
	} else {
		message := tgbotapi.NewMessage(p.chatID, text)
		message.ParseMode = "Markdown"
		_, err = p.bot.Send(message)
	}

	if err != nil {
		return fmt.Errorf("can't send telegram message: %w", err)
	}

	return nil
}
