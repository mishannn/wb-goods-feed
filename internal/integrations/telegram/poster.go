package telegram

import (
	"fmt"

	"example.org/wbsniper/internal/entities/feed"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Poster struct {
	bot         *tgbotapi.BotAPI
	channelName string
}

func NewPoster(botToken string, channelName string) (*Poster, error) {
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		return nil, fmt.Errorf("can't initialize telegram bot: %w", err)
	}

	return &Poster{
		bot:         bot,
		channelName: channelName,
	}, nil
}

func (p *Poster) PublishPost(post feed.Post) error {
	message := tgbotapi.NewMessageToChannel(p.channelName, fmt.Sprintf("*%s*\n\n%s\n\n[Открыть карточку товара](%s)", post.Title, post.Content, post.Link))
	message.ParseMode = "Markdown"

	_, err := p.bot.Send(message)
	if err != nil {
		return fmt.Errorf("can't send telegram message: %w", err)
	}

	return nil
}
