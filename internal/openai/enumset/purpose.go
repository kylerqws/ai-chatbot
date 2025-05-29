package enumset

import base "github.com/kylerqws/chatbot/pkg/openai/domain/purpose"

type PurposeManager struct {
	List map[string]*base.Purpose
}

func NewPurposeManager() *PurposeManager {
	return &PurposeManager{List: base.AllPurposes}
}

func (*PurposeManager) Resolve(code string) (*base.Purpose, error) {
	return base.Resolve(code)
}

func (*PurposeManager) JoinCodes(sep string) string {
	return base.JoinCodes(sep)
}

func (*PurposeManager) Default() *base.Purpose {
	return base.FineTune
}
