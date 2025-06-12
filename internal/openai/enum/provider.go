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

// New returns a new implementation of EnumProvider.
func New() ctrprv.EnumProvider {
	return &provider{}
}

// ChatRole returns an instance of the chat role manager.
func (m *provider) ChatRole() *chatrole.Manager {
	m.chatRoleOnce.Do(func() {
		m.chatRole = chatrole.NewManager()
	})
	return m.chatRole
}

// EventLevel returns an instance of the event level manager.
func (m *provider) EventLevel() *eventlevel.Manager {
	m.eventLevelOnce.Do(func() {
		m.eventLevel = eventlevel.NewManager()
	})
	return m.eventLevel
}

// JobStatus returns an instance of the job status manager.
func (m *provider) JobStatus() *jobstatus.Manager {
	m.jobStatusOnce.Do(func() {
		m.jobStatus = jobstatus.NewManager()
	})
	return m.jobStatus
}

// Model returns an instance of the model manager.
func (m *provider) Model() *model.Manager {
	m.modelOnce.Do(func() {
		m.model = model.NewManager()
	})
	return m.model
}

// Owner returns an instance of the owner manager.
func (m *provider) Owner() *owner.Manager {
	m.ownerOnce.Do(func() {
		m.owner = owner.NewManager()
	})
	return m.owner
}

// Purpose returns an instance of the purpose manager.
func (m *provider) Purpose() *purpose.Manager {
	m.purposeOnce.Do(func() {
		m.purpose = purpose.NewManager()
	})
	return m.purpose
}
