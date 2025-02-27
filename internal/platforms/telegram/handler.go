package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kylerqws/chatgpt-bot/internal/bot"
)

type TelegramHandler struct {
	bot     *TelegramBot
	handler *bot.MessageHandler
}

func NewTelegramHandler(bot *TelegramBot, handler *bot.MessageHandler) *TelegramHandler {
	return &TelegramHandler{
		bot:     bot,
		handler: handler,
	}
}

func (h *TelegramHandler) processMessage(message *tgbotapi.Message) {
	log.Printf("New message from %d: %s", message.From.ID, message.Text)
	h.handler.HandleMessage(h.bot, int(message.From.ID), message.Text)
}
