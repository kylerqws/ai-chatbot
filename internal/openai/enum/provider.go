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

// NewProvider creates a new enum provider that groups OpenAI enum managers.
func NewProvider() ctrprv.EnumProvider {
	return &provider{}
}

// ChatRole returns the enum manager for chat roles.
func (p *provider) ChatRole() *chatrole.Manager {
	p.chatRoleOnce.Do(func() {
		p.chatRole = chatrole.NewManager()
	})
	return p.chatRole
}

// EventLevel returns the enum manager for event levels.
func (p *provider) EventLevel() *eventlevel.Manager {
	p.eventLevelOnce.Do(func() {
		p.eventLevel = eventlevel.NewManager()
	})
	return p.eventLevel
}

// JobStatus returns the enum manager for fine-tuning job statuses.
func (p *provider) JobStatus() *jobstatus.Manager {
	p.jobStatusOnce.Do(func() {
		p.jobStatus = jobstatus.NewManager()
	})
	return p.jobStatus
}

// Model returns the enum manager for models.
func (p *provider) Model() *model.Manager {
	p.modelOnce.Do(func() {
		p.model = model.NewManager()
	})
	return p.model
}

// Owner returns the enum manager for owners.
func (p *provider) Owner() *owner.Manager {
	p.ownerOnce.Do(func() {
		p.owner = owner.NewManager()
	})
	return p.owner
}

// Purpose returns the enum manager for purposes.
func (p *provider) Purpose() *purpose.Manager {
	p.purposeOnce.Do(func() {
		p.purpose = purpose.NewManager()
	})
	return p.purpose
}
