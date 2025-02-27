package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramBot struct {
	bot     *tgbotapi.BotAPI
	handler *TelegramHandler
}

func NewTelegramBot(token string) *TelegramBot {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatalf("Error creating Telegram bot: %v", err)
	}
	return &TelegramBot{bot: bot}
}

func (t *TelegramBot) SetHandler(handler *TelegramHandler) {
	t.handler = handler
}

func (t *TelegramBot) Start() error {
	log.Println("Telegram Bot started")

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := t.bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message != nil {
			go t.handler.processMessage(update.Message)
		}
	}
	return nil
}

func (t *TelegramBot) SendMessage(userID int, message string) error {
	msg := tgbotapi.NewMessage(int64(userID), message)
	_, err := t.bot.Send(msg)
	return err
}
