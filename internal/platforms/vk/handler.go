package vk

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/kylerqws/chatgpt-bot/internal/bot"
)

type VKHandler struct {
	bot     *VKBot
	handler *bot.MessageHandler
}

func NewVKHandler(bot *VKBot, handler *bot.MessageHandler) *VKHandler {
	return &VKHandler{
		bot:     bot,
		handler: handler,
	}
}

func (h *VKHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var request map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	eventType, ok := request["type"].(string)
	if !ok {
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	if eventType == "confirmation" {
		fmt.Fprint(w, os.Getenv("VK_CONFIRMATION_CODE"))
		return
	}

	if eventType == "message_new" {
		go h.processMessage(request)
	}

	fmt.Fprint(w, "ok")
}

func (h *VKHandler) processMessage(request map[string]interface{}) {
	object, ok := request["object"].(map[string]interface{})
	if !ok {
		log.Println("Invalid object format")
		return
	}

	message, ok := object["message"].(map[string]interface{})
	if !ok {
		log.Println("Invalid message format")
		return
	}

	fromID, ok := message["from_id"].(float64)
	if !ok {
		log.Println("Invalid from_id format")
		return
	}

	text, ok := message["text"].(string)
	if !ok {
		log.Println("Invalid text format")
		return
	}

	h.handler.HandleMessage(h.bot, int(fromID), text)
}
