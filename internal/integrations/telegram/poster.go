package telegram

import (
	"fmt"
	"os"

	"example.org/wbsniper/internal/entities/feed"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Poster struct {
	bot    *tgbotapi.BotAPI
	chatID int64
}

func NewPoster(botTokenEnv string, chatID int64) (*Poster, error) {
	// Classic bot token
	botToken := os.Getenv(botTokenEnv)
	if botToken == "" {
		return nil, fmt.Errorf("Bot token env var %s is not set", botTokenEnv)
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
		photos := make([]interface{}, 0, len(post.Images))
		for imageIndex, image := range post.Images[0:5] {
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
