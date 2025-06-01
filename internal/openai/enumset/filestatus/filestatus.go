package filestatus

import base "github.com/kylerqws/chatbot/pkg/openai/domain/filestatus"

type Codes struct {
	Uploaded  string `json:"uploaded"`
	Processed string `json:"processed"`
	Deleted   string `json:"deleted"`
	Error     string `json:"error"`
}

type Manager struct {
	List  map[string]*base.FileStatus
	Codes *Codes
}

func NewManager() *Manager {
	return &Manager{
		List: base.AllFileStatuses,
		Codes: &Codes{
			Uploaded:  base.UploadedCode,
			Processed: base.ProcessedCode,
			Deleted:   base.DeletedCode,
			Error:     base.ErrorCode,
		},
	}
}

func (*Manager) Resolve(code string) (*base.FileStatus, error) {
	return base.Resolve(code)
}

func (*Manager) JoinCodes(sep string) string {
	return base.JoinCodes(sep)
}
