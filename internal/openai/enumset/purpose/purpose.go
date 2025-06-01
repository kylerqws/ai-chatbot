package purpose

import base "github.com/kylerqws/chatbot/pkg/openai/domain/purpose"

type Codes struct {
	FineTune   string `json:"fine_tune"`
	Assistants string `json:"assistants"`
	Batch      string `json:"batch"`
	UserData   string `json:"user_data"`
	Vision     string `json:"vision"`
	Evals      string `json:"evals"`
}

type Manager struct {
	List  map[string]*base.Purpose
	Codes *Codes
}

func NewManager() *Manager {
	return &Manager{
		List: base.AllPurposes,
		Codes: &Codes{
			FineTune:   base.FineTuneCode,
			Assistants: base.AssistantsCode,
			Batch:      base.BatchCode,
			UserData:   base.UserDataCode,
			Vision:     base.VisionCode,
			Evals:      base.EvalsCode,
		},
	}
}

func (*Manager) Resolve(code string) (*base.Purpose, error) {
	return base.Resolve(code)
}

func (*Manager) JoinCodes(sep string) string {
	return base.JoinCodes(sep)
}

func (*Manager) Default() *base.Purpose {
	return base.FineTune
}
