package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/joho/godotenv"
	"github.com/maya-florenko/arta/internal/app"
	"github.com/maya-florenko/arta/internal/banner"
	"github.com/maya-florenko/arta/internal/ollama"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	banner.Print()

	_ = godotenv.Load()

	if err := ollama.New(); err != nil {
		log.Fatal(err)
	}

	if err := app.Init(ctx); err != nil {
		log.Fatal(err)
	}
}
