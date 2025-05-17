package enumset

import base "github.com/kylerqws/chatbot/pkg/openai/domain/purpose"

type PurposeManager struct {
	List map[string]*base.Purpose
}

func NewPurposeManager() *PurposeManager {
	return &PurposeManager{List: base.AllPurposes}
}

func (_ *PurposeManager) Resolve(code string) (*base.Purpose, error) {
	return base.Resolve(code)
}

func (_ *PurposeManager) JoinCodes(sep string) string {
	return base.JoinCodes(sep)
}
