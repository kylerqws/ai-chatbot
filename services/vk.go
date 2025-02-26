package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/kylerqws/go-chatgpt-vk/handlers"
)

type VKService struct {
	Token            string
	ConfirmationCode string
	VKClient         *api.VK
	GetAIResponse    handlers.AIResponseFunc
}

func NewVKService(getAIResponse handlers.AIResponseFunc) *VKService {
	token := os.Getenv("VK_TOKEN")
	confirmationCode := os.Getenv("VK_CONFIRMATION_CODE")

	return &VKService{
		Token:            token,
		ConfirmationCode: confirmationCode,
		VKClient:         api.NewVK(token),
		GetAIResponse:    getAIResponse,
	}
}

func (v *VKService) HandleRequest(w http.ResponseWriter, r *http.Request) {
	var request map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Printf("Error decoding JSON: %v", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	if request["type"] == "confirmation" {
		fmt.Fprint(w, v.ConfirmationCode)
		return
	}

	if request["type"] == "message_new" {
		go v.handleMessage(request)
	}

	fmt.Fprint(w, "ok")
	log.Println("Response sent: ok")
}

func (v *VKService) handleMessage(request map[string]interface{}) {
	object := request["object"].(map[string]interface{})
	message := object["message"].(map[string]interface{})
	fromID := int(message["from_id"].(float64))
	userMessage := message["text"].(string)

	log.Printf("New message from user %d: %s", fromID, userMessage)

	aiResponse, err := v.GetAIResponse(userMessage)
	if err != nil {
		log.Printf("Error getting AI response: %v", err)
		aiResponse = "Sorry, I can't respond right now. Please try again later."
	}

	_, err = v.VKClient.MessagesSend(api.Params{
		"user_id":   fromID,
		"random_id": 0,
		"message":   aiResponse,
	})
	if err != nil {
		log.Printf("Error sending message: %v", err)
	} else {
		log.Printf("Message sent to user %d.", fromID)
	}
}

func (v *VKService) GetName() string {
	return "vk"
}
