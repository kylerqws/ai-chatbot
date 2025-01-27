package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/kylerqws/go-chatgpt-vk.git/services"
)

func main() {
	// Load environment variables from .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Load environment variables
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	// Route requests based on service
	http.HandleFunc("/vk", services.NewVKService().HandleRequest)

	log.Printf("Server started on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
