package handlers

import (
	"log"
	"os"
)

type AIResponseFunc func(string, string) (string, error)

func GetAIResponseHandler() AIResponseFunc {
	aiProvider := os.Getenv("AI_PROVIDER")

	switch aiProvider {
	case "openai":
		log.Println("Using OpenAI API provider")
		return GetOpenAIResponse
	default:
		log.Fatalf("Unsupported AI provider: %s", aiProvider)
		return nil
	}
}
