package jobstatus

import base "github.com/kylerqws/chatbot/pkg/openai/enumset/jobstatus"

// Codes defines available job status codes.
type Codes struct {
	ValidatingFiles string `json:"validating_files"`
	Running         string `json:"running"`
	Succeeded       string `json:"succeeded"`
	Cancelled       string `json:"cancelled"`
	Failed          string `json:"failed"`
}

// Manager provides access to available job status values.
type Manager struct {
	List  map[string]*base.JobStatus
	Codes *Codes
}

// NewManager creates a new manager for available job status values.
func NewManager() *Manager {
	return &Manager{
		List: base.AllJobStatuses,
		Codes: &Codes{
			ValidatingFiles: base.ValidatingFilesCode,
			Running:         base.RunningCode,
			Succeeded:       base.SucceededCode,
			Cancelled:       base.CancelledCode,
			Failed:          base.FailedCode,
		},
	}
}

// Resolve returns the job status associated with the given code.
func (*Manager) Resolve(code string) (*base.JobStatus, error) {
	return base.Resolve(code)
}

// JoinCodes joins all known job status codes using the given separator.
func (*Manager) JoinCodes(sep string) string {
	return base.JoinCodes(sep)
}
