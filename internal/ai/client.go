package ai

import (
	"fmt"
	"os"
)

// AIClient - интерфейс для всех AI-сервисов (OpenAI, Azure, Claude и т. д.).
type AIClient interface {
	GetResponse(prompt string) (string, error)
}

// NewAIClient создаёт AI клиента в зависимости от сервиса (OpenAI, Azure и т. д.).
func NewAIClient() (AIClient, error) {
	provider := os.Getenv("AI_PROVIDER") // openai, azure, claude и т. д.

	switch provider {
	case "openai":
		apiKey := os.Getenv("OPENAI_API_KEY")
		if apiKey == "" {
			return nil, fmt.Errorf("missing OPENAI_API_KEY")
		}
		return NewOpenAIClient(apiKey), nil

	default:
		return nil, fmt.Errorf("unsupported AI provider: %s", provider)
	}
}
