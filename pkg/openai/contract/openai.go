package contract

import ctrsvc "github.com/kylerqws/chatbot/pkg/openai/contract/service"

// OpenAI aggregates access to OpenAI API services.
type OpenAI interface {
	// FileService handles file uploads and management.
	FileService() ctrsvc.FileService

	// FineTuningService handles fine-tuning jobs and results.
	FineTuningService() ctrsvc.FineTuningService

	// ModelService manages available models.
	ModelService() ctrsvc.ModelService

	// ChatService interacts with chat-based language models.
	ChatService() ctrsvc.ChatService
}
