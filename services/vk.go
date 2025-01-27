package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/kylerqws/go-chatgpt-vk.git/handlers"
)

type VKService struct {
	Token            string
	ConfirmationCode string
	VKClient         *api.VK
}

func NewVKService() *VKService {
	token := os.Getenv("VK_TOKEN")
	confirmationCode := os.Getenv("VK_CONFIRMATION_CODE")
	return &VKService{
		Token:            token,
		ConfirmationCode: confirmationCode,
		VKClient:         api.NewVK(token),
	}
}

func (v *VKService) HandleRequest(w http.ResponseWriter, r *http.Request) {
	var request map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Printf("error decoding JSON: %v", err)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	// Handle confirmation
	if request["type"] == "confirmation" {
		fmt.Fprint(w, v.ConfirmationCode)
		return
	}

	// Handle new messages
	if request["type"] == "message_new" {
		go v.handleMessage(request)
	}

	// Response for VK
	fmt.Fprint(w, "ok")
	log.Println("response sent: ok")
}

func (v *VKService) handleMessage(request map[string]interface{}) {
	object := request["object"].(map[string]interface{})
	message := object["message"].(map[string]interface{})
	fromID := int(message["from_id"].(float64))
	userMessage := message["text"].(string)

	// Log incoming message
	log.Printf("new message from user %d: %s", fromID, userMessage)

	// Get response from ChatGPT
	chatGPTResponse, err := handlers.GetChatGPTResponse(userMessage, os.Getenv("OPENAI_API_KEY"))
	if err != nil {
		log.Printf("error getting response from ChatGPT: %v", err)
		chatGPTResponse = "Sorry, I can't respond right now. Please try again later."
	}

	// Send response to user
	_, err = v.VKClient.MessagesSend(api.Params{
		"user_id":   fromID,
		"random_id": 0,
		"message":   chatGPTResponse,
	})
	if err != nil {
		log.Printf("error sending message: %v", err)
	} else {
		log.Printf("message sent to user %d.", fromID)
	}
}

func (v *VKService) GetName() string {
	return "vk"
}
