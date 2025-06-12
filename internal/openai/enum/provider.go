package enum

import (
	"sync"

	"github.com/kylerqws/chatbot/internal/openai/enum/chatrole"
	"github.com/kylerqws/chatbot/internal/openai/enum/eventlevel"
	"github.com/kylerqws/chatbot/internal/openai/enum/jobstatus"
	"github.com/kylerqws/chatbot/internal/openai/enum/model"
	"github.com/kylerqws/chatbot/internal/openai/enum/owner"
	"github.com/kylerqws/chatbot/internal/openai/enum/purpose"

	ctrprv "github.com/kylerqws/chatbot/internal/openai/contract/provider"
)

// provider provides access to OpenAI enum managers.
type provider struct {
	chatRoleOnce sync.Once
	chatRole     *chatrole.Manager

	eventLevelOnce sync.Once
	eventLevel     *eventlevel.Manager

	jobStatusOnce sync.Once
	jobStatus     *jobstatus.Manager

	modelOnce sync.Once
	model     *model.Manager

	ownerOnce sync.Once
	owner     *owner.Manager

	purposeOnce sync.Once
	purpose     *purpose.Manager
}

// NewProvider returns a new enum provider that groups OpenAI enum managers.
func NewProvider() ctrprv.EnumProvider {
	return &provider{}
}

// ChatRole returns the enum manager for chat roles.
func (m *provider) ChatRole() *chatrole.Manager {
	m.chatRoleOnce.Do(func() {
		m.chatRole = chatrole.NewManager()
	})
	return m.chatRole
}

// EventLevel returns the enum manager for event levels.
func (m *provider) EventLevel() *eventlevel.Manager {
	m.eventLevelOnce.Do(func() {
		m.eventLevel = eventlevel.NewManager()
	})
	return m.eventLevel
}

// JobStatus returns the enum manager for fine-tuning job statuses.
func (m *provider) JobStatus() *jobstatus.Manager {
	m.jobStatusOnce.Do(func() {
		m.jobStatus = jobstatus.NewManager()
	})
	return m.jobStatus
}

// Model returns the enum manager for models.
func (m *provider) Model() *model.Manager {
	m.modelOnce.Do(func() {
		m.model = model.NewManager()
	})
	return m.model
}

// Owner returns the enum manager for owners.
func (m *provider) Owner() *owner.Manager {
	m.ownerOnce.Do(func() {
		m.owner = owner.NewManager()
	})
	return m.owner
}

// Purpose returns the enum manager for file purposes.
func (m *provider) Purpose() *purpose.Manager {
	m.purposeOnce.Do(func() {
		m.purpose = purpose.NewManager()
	})
	return m.purpose
}
