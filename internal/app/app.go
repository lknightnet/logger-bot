package app

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"log"
	"logger-bot/config"
	"logger-bot/internal/controller"
	"logger-bot/internal/service"
	"os"
	"os/signal"
)

func Run(cfg *config.Config) {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithMiddlewares(showMessageWithUserName),
	}

	b, err := bot.New(cfg.TG.Token, opts...)
	if err != nil {
		log.Println(err)
	}

	ms := service.NewMessageService(cfg.API.UrlAPI)

	controller.NewControllers(b, ms)

	b.Start(ctx)
}

func showMessageWithUserName(next bot.HandlerFunc) bot.HandlerFunc {
	return func(ctx context.Context, b *bot.Bot, update *models.Update) {
		if update.Message != nil {
			log.Printf("%s say: %s", update.Message.From.FirstName, update.Message.Text)
		}
		next(ctx, b, update)
	}
}
