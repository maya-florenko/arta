package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/joho/godotenv"
	"github.com/mayusha256/arta/internal/app"
	"github.com/mayusha256/arta/internal/banner"
)

func main() {
	c, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	banner.Print()

	app.Run(c)
}
