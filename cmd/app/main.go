package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/mayusha256/arta/internal/app"
	"github.com/mayusha256/arta/internal/banner"
)

func main() {
	c, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	banner.Print()

	app.Run(c)
}
