package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kylerqws/chatgpt-bot/cmd"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found, using system environment variables")
	}

	cmd.Execute()
}
