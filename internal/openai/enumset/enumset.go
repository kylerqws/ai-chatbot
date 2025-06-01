package enumset

import (
	"github.com/kylerqws/chatbot/internal/openai/enumset/chatrole"
	"github.com/kylerqws/chatbot/internal/openai/enumset/jobstatus"
	"github.com/kylerqws/chatbot/internal/openai/enumset/model"
	"github.com/kylerqws/chatbot/internal/openai/enumset/purpose"
)

type ManagerSet struct {
	chatRole  *chatrole.Manager
	purpose   *purpose.Manager
	model     *model.Manager
	jobStatus *jobstatus.Manager
}

func NewManagerSet() *ManagerSet {
	return &ManagerSet{}
}

func (s *ManagerSet) ChatRole() *chatrole.Manager {
	if s.chatRole == nil {
		s.chatRole = chatrole.NewManager()
	}
	return s.chatRole
}

func (s *ManagerSet) Purpose() *purpose.Manager {
	if s.purpose == nil {
		s.purpose = purpose.NewManager()
	}
	return s.purpose
}

func (s *ManagerSet) Model() *model.Manager {
	if s.model == nil {
		s.model = model.NewManager()
	}
	return s.model
}

func (s *ManagerSet) JobStatus() *jobstatus.Manager {
	if s.jobStatus == nil {
		s.jobStatus = jobstatus.NewManager()
	}
	return s.jobStatus
}
