package controller

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"log"
	"logger-bot/internal/service"
)

type messageController struct {
	ML service.MethodLogger
}

func newMessageController(ml service.MethodLogger) *messageController {
	return &messageController{ML: ml}
}

func NewControllers(b *bot.Bot, ml service.MethodLogger) {
	ctrl := newMessageController(ml)

	b.RegisterHandler(bot.HandlerTypeMessageText, "/logs", bot.MatchTypeExact, ctrl.logs)
}

func (mc *messageController) logs(ctx context.Context, b *bot.Bot, update *models.Update) {
	msgs, err := mc.ML.GetLogs()
	if err != nil {
		log.Println(err)
	}
	for _, value := range msgs {

		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   value.String(),
		})
	}
}
