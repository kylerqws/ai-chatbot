package vk

import (
	"log"

	"github.com/SevereCloud/vksdk/v2/api"
)

type VKBot struct {
	client *api.VK
}

func NewVKBot(token string) *VKBot {
	return &VKBot{
		client: api.NewVK(token),
	}
}

func (v *VKBot) Start() error {
	log.Println("VK Bot started")
	return nil
}

func (v *VKBot) SendMessage(userID int, message string) error {
	_, err := v.client.MessagesSend(api.Params{
		"user_id":   userID,
		"random_id": 0,
		"message":   message,
	})
	return err
}
