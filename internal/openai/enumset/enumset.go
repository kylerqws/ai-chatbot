package enumset

import (
	"sync"

	setchr "github.com/kylerqws/chatbot/internal/openai/enumset/chatrole"
	setevl "github.com/kylerqws/chatbot/internal/openai/enumset/eventlevel"
	setjst "github.com/kylerqws/chatbot/internal/openai/enumset/jobstatus"
	setmdl "github.com/kylerqws/chatbot/internal/openai/enumset/model"
	setown "github.com/kylerqws/chatbot/internal/openai/enumset/owner"
	setprp "github.com/kylerqws/chatbot/internal/openai/enumset/purpose"
)

// ManagerSet provides access to enumset managers.
type ManagerSet struct {
	chatOnce sync.Once
	chatRole *setchr.Manager

	eventOnce  sync.Once
	eventLevel *setevl.Manager

	jobOnce   sync.Once
	jobStatus *setjst.Manager

	modelOnce sync.Once
	model     *setmdl.Manager

	ownerOnce sync.Once
	owner     *setown.Manager

	purposeOnce sync.Once
	purpose     *setprp.Manager
}

// NewManagerSet returns a new ManagerSet.
func NewManagerSet() *ManagerSet {
	return &ManagerSet{}
}

// ChatRole returns the chat role manager.
func (s *ManagerSet) ChatRole() *setchr.Manager {
	s.chatOnce.Do(func() {
		s.chatRole = setchr.NewManager()
	})
	return s.chatRole
}

// EventLevel returns the event level manager.
func (s *ManagerSet) EventLevel() *setevl.Manager {
	s.eventOnce.Do(func() {
		s.eventLevel = setevl.NewManager()
	})
	return s.eventLevel
}

// JobStatus returns the job status manager.
func (s *ManagerSet) JobStatus() *setjst.Manager {
	s.jobOnce.Do(func() {
		s.jobStatus = setjst.NewManager()
	})
	return s.jobStatus
}

// Model returns the model manager.
func (s *ManagerSet) Model() *setmdl.Manager {
	s.modelOnce.Do(func() {
		s.model = setmdl.NewManager()
	})
	return s.model
}

// Owner returns the owner manager.
func (s *ManagerSet) Owner() *setown.Manager {
	s.ownerOnce.Do(func() {
		s.owner = setown.NewManager()
	})
	return s.owner
}

// Purpose returns the purpose manager.
func (s *ManagerSet) Purpose() *setprp.Manager {
	s.purposeOnce.Do(func() {
		s.purpose = setprp.NewManager()
	})
	return s.purpose
}
