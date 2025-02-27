package bot

import (
	"log"
	"sync"

	"github.com/kylerqws/chatgpt-bot/internal/database"
)

// MessageHandler - обработчик сообщений.
type MessageHandler struct {
	mu       sync.Mutex
	aiClient AIClient
	db       database.DatabaseClient
}

// AIClient - интерфейс для работы с AI.
type AIClient interface {
	GetResponse(prompt string) (string, error)
}

// NewMessageHandler создаёт новый обработчик сообщений.
func NewMessageHandler(aiClient AIClient, db database.DatabaseClient) *MessageHandler {
	return &MessageHandler{aiClient: aiClient, db: db}
}

// HandleMessage обрабатывает сообщение и сохраняет в БД.
func (h *MessageHandler) HandleMessage(bot Client, userID int, message string) {
	h.mu.Lock()
	defer h.mu.Unlock()

	log.Printf("Received message from user %d: %s", userID, message)

	response, err := h.aiClient.GetResponse(message)
	if err != nil {
		log.Printf("Error getting AI response: %v", err)
		response = "An error occurred. Please try again later."
	}

	if err := bot.SendMessage(userID, response); err != nil {
		log.Printf("Error sending message: %v", err)
	}

	// Сохраняем сообщение в БД
	err = h.db.SaveMessage(userID, message, response)
	if err != nil {
		log.Printf("Error saving message: %v", err)
	}
}
