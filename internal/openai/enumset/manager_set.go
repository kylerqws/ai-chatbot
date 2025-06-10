package enumset

import (
	"sync"

	"github.com/kylerqws/chatbot/internal/openai/enumset/chatrole"
	"github.com/kylerqws/chatbot/internal/openai/enumset/eventlevel"
	"github.com/kylerqws/chatbot/internal/openai/enumset/jobstatus"
	"github.com/kylerqws/chatbot/internal/openai/enumset/model"
	"github.com/kylerqws/chatbot/internal/openai/enumset/owner"
	"github.com/kylerqws/chatbot/internal/openai/enumset/purpose"

	ctrenm "github.com/kylerqws/chatbot/internal/openai/contract/enumset"
)

// managerSet provides thread-safe access to OpenAI enum managers.
type managerSet struct {
	chatOnce sync.Once
	chatRole *chatrole.Manager

	eventOnce  sync.Once
	eventLevel *eventlevel.Manager

	jobOnce   sync.Once
	jobStatus *jobstatus.Manager

	modelOnce sync.Once
	model     *model.Manager

	ownerOnce sync.Once
	owner     *owner.Manager

	purposeOnce sync.Once
	purpose     *purpose.Manager
}

// NewManagerSet returns a new enum manager set.
func NewManagerSet() ctrenm.ManagerSet {
	return &managerSet{}
}

// ChatRole returns the chat role manager.
func (m *managerSet) ChatRole() *chatrole.Manager {
	m.chatOnce.Do(func() {
		m.chatRole = chatrole.NewManager()
	})
	return m.chatRole
}

// EventLevel returns the event level manager.
func (m *managerSet) EventLevel() *eventlevel.Manager {
	m.eventOnce.Do(func() {
		m.eventLevel = eventlevel.NewManager()
	})
	return m.eventLevel
}

// JobStatus returns the job status manager.
func (m *managerSet) JobStatus() *jobstatus.Manager {
	m.jobOnce.Do(func() {
		m.jobStatus = jobstatus.NewManager()
	})
	return m.jobStatus
}

// Model returns the model manager.
func (m *managerSet) Model() *model.Manager {
	m.modelOnce.Do(func() {
		m.model = model.NewManager()
	})
	return m.model
}

// Owner returns the owner manager.
func (m *managerSet) Owner() *owner.Manager {
	m.ownerOnce.Do(func() {
		m.owner = owner.NewManager()
	})
	return m.owner
}

// Purpose returns the purpose manager.
func (m *managerSet) Purpose() *purpose.Manager {
	m.purposeOnce.Do(func() {
		m.purpose = purpose.NewManager()
	})
	return m.purpose
}
