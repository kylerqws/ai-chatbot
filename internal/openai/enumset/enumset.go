package enumset

import (
	"github.com/kylerqws/chatbot/internal/openai/enumset/chatrole"
	"github.com/kylerqws/chatbot/internal/openai/enumset/filestatus"
	"github.com/kylerqws/chatbot/internal/openai/enumset/jobstatus"
	"github.com/kylerqws/chatbot/internal/openai/enumset/model"
	"github.com/kylerqws/chatbot/internal/openai/enumset/purpose"
)

type ManagerSet struct {
	chatRole   *chatrole.Manager
	fileStatus *filestatus.Manager
	jobStatus  *jobstatus.Manager
	model      *model.Manager
	purpose    *purpose.Manager
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

func (s *ManagerSet) FileStatus() *filestatus.Manager {
	if s.fileStatus == nil {
		s.fileStatus = filestatus.NewManager()
	}
	return s.fileStatus
}

func (s *ManagerSet) JobStatus() *jobstatus.Manager {
	if s.jobStatus == nil {
		s.jobStatus = jobstatus.NewManager()
	}
	return s.jobStatus
}

func (s *ManagerSet) Model() *model.Manager {
	if s.model == nil {
		s.model = model.NewManager()
	}
	return s.model
}

func (s *ManagerSet) Purpose() *purpose.Manager {
	if s.purpose == nil {
		s.purpose = purpose.NewManager()
	}
	return s.purpose
}
