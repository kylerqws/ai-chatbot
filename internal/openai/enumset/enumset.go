package enumset

import (
	"github.com/kylerqws/chatbot/internal/openai/enumset/jobstatus"
	"github.com/kylerqws/chatbot/internal/openai/enumset/model"
	"github.com/kylerqws/chatbot/internal/openai/enumset/purpose"
)

type Set struct {
	purposeManager   *purpose.Manager
	modelManager     *model.Manager
	jobStatusManager *jobstatus.Manager
}

func NewSet() *Set {
	return &Set{}
}

func (s *Set) PurposeManager() *purpose.Manager {
	if s.purposeManager == nil {
		s.purposeManager = purpose.NewManager()
	}
	return s.purposeManager
}

func (s *Set) ModelManager() *model.Manager {
	if s.modelManager == nil {
		s.modelManager = model.NewManager()
	}
	return s.modelManager
}

func (s *Set) JobStatusManager() *jobstatus.Manager {
	if s.jobStatusManager == nil {
		s.jobStatusManager = jobstatus.NewManager()
	}
	return s.jobStatusManager
}
