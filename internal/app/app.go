package app

import (
	"context"
	"log"
	"os"

	"github.com/go-telegram/bot"
)

func Run(c context.Context) {
	opts := []bot.Option{
		bot.WithDefaultHandler(Handler),
	}

	b, err := bot.New(os.Getenv("TELEGRAM_TOKEN"), opts...)
	if err != nil {
		log.Fatal(err)
	}

	b.Start(c)
}
