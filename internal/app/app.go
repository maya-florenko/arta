package app

import (
	"context"
	"os"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/maya-florenko/arta/internal/ollama"
)

func Init(ctx context.Context) error {
	opts := []bot.Option{
		bot.WithDefaultHandler(handler),
	}

	b, err := bot.New(os.Getenv("TELEGRAM_TOKEN"), opts...)
	if err != nil {
		return err
	}

	b.Start(ctx)

	return nil
}

func handler(ctx context.Context, b *bot.Bot, u *models.Update) {
	if u.Message == nil {
		return
	}

	b.SendChatAction(ctx, &bot.SendChatActionParams{
		ChatID: u.Message.Chat.ID,
		Action: models.ChatActionTyping,
	})

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:    u.Message.Chat.ID,
		Text:      ollama.Generate(ctx, "user "+u.Message.From.FirstName+" is write you: "+u.Message.Text),
		ParseMode: models.ParseModeHTML,
	})
}
