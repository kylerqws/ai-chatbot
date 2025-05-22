package enumset

import base "github.com/kylerqws/chatbot/pkg/openai/domain/jobstatus"

type JobStatusManager struct {
	List map[string]*base.JobStatus
}

func NewJobStatusManager() *JobStatusManager {
	return &JobStatusManager{List: base.AllJobStatuses}
}

func (*JobStatusManager) Resolve(code string) (*base.JobStatus, error) {
	return base.Resolve(code)
}

func (*JobStatusManager) JoinCodes(sep string) string {
	return base.JoinCodes(sep)
}
