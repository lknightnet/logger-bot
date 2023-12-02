package main

import (
	"log"
	"logger-bot/config"
	"logger-bot/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Println(err)
	}

	app.Run(cfg)
}
