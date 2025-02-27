package cmd

import (
	"log"
	"net/http"
	"os"

	"github.com/kylerqws/chatgpt-bot/internal/ai"
	"github.com/kylerqws/chatgpt-bot/internal/bot"
	"github.com/kylerqws/chatgpt-bot/internal/database"
	"github.com/kylerqws/chatgpt-bot/internal/platforms/telegram"
	"github.com/kylerqws/chatgpt-bot/internal/platforms/vk"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start chat bot server",
	Run: func(cmd *cobra.Command, args []string) {
		startServer()
	},
}

func startServer() {
	// Подключаем БД через фабрику
	dbClient, err := database.NewDatabaseClient()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer dbClient.Close()

	aiClient, err := ai.NewAIClient()
	if err != nil {
		log.Fatalf("Failed to initialize AI client: %v", err)
	}

	messageHandler := bot.NewMessageHandler(aiClient, dbClient)

	var bots []bot.Client

	vkToken := os.Getenv("VK_TOKEN")
	if vkToken != "" {
		log.Println("VK Bot is enabled")
		vkBot := vk.NewVKBot(vkToken)
		vkHandler := vk.NewVKHandler(vkBot, messageHandler)
		bots = append(bots, vkBot)
		go startVKServer(vkHandler)
	}

	telegramToken := os.Getenv("TELEGRAM_TOKEN")
	if telegramToken != "" {
		log.Println("Telegram Bot is enabled")
		telegramBot := telegram.NewTelegramBot(telegramToken)
		telegramHandler := telegram.NewTelegramHandler(telegramBot, messageHandler)
		telegramBot.SetHandler(telegramHandler)
		bots = append(bots, telegramBot)
		go telegramBot.Start()
	}

	if len(bots) == 0 {
		log.Fatal("Error: No bot is configured (VK_TOKEN or TELEGRAM_TOKEN is missing)")
	}

	log.Println("All bots started successfully")
	select {}
}

func startVKServer(handler *vk.VKHandler) {
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "5000"
	}
	log.Printf("VK server is listening on port %s...", port)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatalf("Error starting VK server: %v", err)
	}
}
