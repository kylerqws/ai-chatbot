package enumset

import base "github.com/kylerqws/chatbot/pkg/openai/domain/model"

type ModelManager struct {
	List map[string]*base.Model
}

func NewModelManager() *ModelManager {
	return &ModelManager{List: base.AllModels}
}

func (*ModelManager) Resolve(code string) (*base.Model, error) {
	return base.Resolve(code)
}

func (*ModelManager) JoinCodes(sep string) string {
	return base.JoinCodes(sep)
}
