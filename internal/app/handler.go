package app

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/mayusha256/arta/internal/ai"
)

func Handler(ctx context.Context, b *bot.Bot, u *models.Update) {
	if u.Message == nil {
		return
	}

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: u.Message.Chat.ID,
		Text:   ai.Generate(ctx, u.Message.Text),
	})
}
