package contract

import ctrprv "github.com/kylerqws/chatbot/internal/openai/contract/provider"

// OpenAI defines the root contract for accessing OpenAI service and enum providers.
type OpenAI interface {
	// ServiceProvider returns grouped OpenAI API services.
	ServiceProvider() ctrprv.ServiceProvider

	// EnumProvider returns grouped enum managers used in OpenAI operations.
	EnumProvider() ctrprv.EnumProvider
}
