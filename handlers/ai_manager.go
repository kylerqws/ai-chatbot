package handlers

import (
	"log"
	"os"
)

type AIResponseFunc func(string) (string, error)

func GetAIResponseHandler() AIResponseFunc {
	aiProvider := os.Getenv("AI_PROVIDER")
	apiKey := os.Getenv("OPENAI_API_KEY")

	switch aiProvider {
	case "openai":
		log.Println("Using OpenAI API provider")
		return func(prompt string) (string, error) {
			return GetOpenAIResponse(prompt, apiKey)
		}
	default:
		log.Fatalf("Unsupported AI provider: %s", aiProvider)
		return nil
	}
}
