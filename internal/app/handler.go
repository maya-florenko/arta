package app

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func Handler(c context.Context, b *bot.Bot, u *models.Update) {
	b.SendMessage(c, &bot.SendMessageParams{
		ChatID: u.Message.Chat.ID,
		Text:   u.Message.Text,
	})
}
