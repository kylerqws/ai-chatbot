package jobstatus

import base "github.com/kylerqws/chatbot/pkg/openai/domain/jobstatus"

type StatusCodes struct {
	Validating string `json:"validating"`
	Running    string `json:"running"`
	Succeeded  string `json:"succeeded"`
	Failed     string `json:"failed"`
	Cancelled  string `json:"cancelled"`
}

type Manager struct {
	List  map[string]*base.JobStatus
	Codes *StatusCodes
}

func NewManager() *Manager {
	return &Manager{
		List: base.AllJobStatuses,
		Codes: &StatusCodes{
			Validating: base.ValidatingCode,
			Running:    base.RunningCode,
			Succeeded:  base.SucceededCode,
			Failed:     base.FailedCode,
			Cancelled:  base.CancelledCode,
		},
	}
}

func (*Manager) Resolve(code string) (*base.JobStatus, error) {
	return base.Resolve(code)
}

func (*Manager) JoinCodes(sep string) string {
	return base.JoinCodes(sep)
}
