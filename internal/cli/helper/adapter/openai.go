package adapter

import (
	"fmt"
	"github.com/spf13/cobra"

	"github.com/kylerqws/chatbot/internal/openai/enumset"
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

func (h *OpenAiAdapter) ValidatePurposeCode(prpCode string) error {
	if _, err := h.prpManager.Resolve(prpCode); err != nil {
		return fmt.Errorf("invalid purpose code %q: %w", prpCode, err)
	}
	return nil
}
