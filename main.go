package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/kylerqws/go-chatgpt-vk/handlers"
	"github.com/kylerqws/go-chatgpt-vk/services"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	vkService := services.NewVKService(handlers.GetAIResponseHandler())

	http.HandleFunc("/vk", vkService.HandleRequest)

	log.Printf("Server started on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
