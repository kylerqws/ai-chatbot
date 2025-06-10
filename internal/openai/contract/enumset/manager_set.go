package enumset

import (
	"github.com/kylerqws/chatbot/internal/openai/enumset/chatrole"
	"github.com/kylerqws/chatbot/internal/openai/enumset/eventlevel"
	"github.com/kylerqws/chatbot/internal/openai/enumset/jobstatus"
	"github.com/kylerqws/chatbot/internal/openai/enumset/model"
	"github.com/kylerqws/chatbot/internal/openai/enumset/owner"
	"github.com/kylerqws/chatbot/internal/openai/enumset/purpose"
)

// ManagerSet defines access to enum managers.
type ManagerSet interface {
	// ChatRole returns the chat role manager.
	ChatRole() *chatrole.Manager

	// EventLevel returns the event level manager.
	EventLevel() *eventlevel.Manager

	// JobStatus returns the job status manager.
	JobStatus() *jobstatus.Manager

	// Model returns the model manager.
	Model() *model.Manager

	// Owner returns the owner manager.
	Owner() *owner.Manager

	// Purpose returns the purpose manager.
	Purpose() *purpose.Manager
}
