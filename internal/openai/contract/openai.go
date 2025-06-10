package contract

import (
	ctrenm "github.com/kylerqws/chatbot/internal/openai/contract/enumset"
	ctrsvc "github.com/kylerqws/chatbot/pkg/openai/contract/service"
)

// OpenAI aggregates access to OpenAI API services and enum managers.
type OpenAI interface {
	// ChatService returns the service for chat interactions.
	ChatService() ctrsvc.ChatService

	// FileService returns the service for file operations.
	FileService() ctrsvc.FileService

	// FineTuningService returns the service for fine-tuning jobs.
	FineTuningService() ctrsvc.FineTuningService

	// ModelService returns the service for model management.
	ModelService() ctrsvc.ModelService

	// ManagerSet returns the manager set for OpenAI enum types.
	ManagerSet() ctrenm.ManagerSet
}
