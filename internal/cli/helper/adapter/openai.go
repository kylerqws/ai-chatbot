package adapter

import (
	"github.com/kylerqws/chatbot/internal/openai/enumset"
	"github.com/spf13/cobra"
)

type OpenAiAdapter struct {
	command    *cobra.Command
	prpManager *enumset.PurposeManager
}

func NewOpenAiAdapter(cmd *cobra.Command) *OpenAiAdapter {
	return &OpenAiAdapter{command: cmd}
}

func (h *OpenAiAdapter) PurposeManager() *enumset.PurposeManager {
	if h.prpManager == nil {
		h.prpManager = enumset.NewPurposeManager()
	}
	return h.prpManager
}
