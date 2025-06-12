package provider

import (
	"github.com/kylerqws/chatbot/internal/openai/enumset/chatrole"
	"github.com/kylerqws/chatbot/internal/openai/enumset/eventlevel"
	"github.com/kylerqws/chatbot/internal/openai/enumset/jobstatus"
	"github.com/kylerqws/chatbot/internal/openai/enumset/model"
	"github.com/kylerqws/chatbot/internal/openai/enumset/owner"
	"github.com/kylerqws/chatbot/internal/openai/enumset/purpose"
)

// EnumProvider defines grouped access to OpenAI enum managers.
type EnumProvider interface {
	// ChatRole returns the enum manager for chat roles.
	ChatRole() *chatrole.Manager

	// EventLevel returns the enum manager for event levels.
	EventLevel() *eventlevel.Manager

	// JobStatus returns the enum manager for fine-tuning job statuses.
	JobStatus() *jobstatus.Manager

	// Model returns the enum manager for models.
	Model() *model.Manager

	// Owner returns the enum manager for model ownership.
	Owner() *owner.Manager

	// Purpose returns the enum manager for file purposes.
	Purpose() *purpose.Manager
}
