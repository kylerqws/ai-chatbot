package provider

import ctrsvc "github.com/kylerqws/chatbot/internal/openai/contract/service"

// ServiceProvider defines grouped access to OpenAI API services.
type ServiceProvider interface {
	// Chat returns the service for chat completions.
	Chat() ctrsvc.ChatService

	// File returns the service for file management.
	File() ctrsvc.FileService

	// FineTuning returns the service for fine-tuning jobs.
	FineTuning() ctrsvc.FineTuningService

	// Model returns the service for model operations.
	Model() ctrsvc.ModelService
}
