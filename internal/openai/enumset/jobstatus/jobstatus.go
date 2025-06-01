package jobstatus

import base "github.com/kylerqws/chatbot/pkg/openai/domain/jobstatus"

type Codes struct {
	ValidatingCode string `json:"validating_code"`
	RunningCode    string `json:"running_code"`
	SucceededCode  string `json:"succeeded_code"`
	FailedCode     string `json:"failed_code"`
	CancelledCode  string `json:"cancelled_code"`
}

type Manager struct {
	List  map[string]*base.JobStatus
	Codes *Codes
}

func NewManager() *Manager {
	return &Manager{List: base.AllJobStatuses, Codes: &Codes{
		ValidatingCode: base.ValidatingCode,
		RunningCode:    base.RunningCode,
		SucceededCode:  base.SucceededCode,
		FailedCode:     base.FailedCode,
		CancelledCode:  base.CancelledCode,
	}}
}

func (*Manager) Resolve(code string) (*base.JobStatus, error) {
	return base.Resolve(code)
}

func (*Manager) JoinCodes(sep string) string {
	return base.JoinCodes(sep)
}
