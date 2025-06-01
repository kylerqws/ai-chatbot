package adapter

import (
	"fmt"
	"github.com/spf13/cobra"

	"github.com/kylerqws/chatbot/internal/openai/enumset"
	"github.com/kylerqws/chatbot/internal/openai/enumset/filestatus"
	"github.com/kylerqws/chatbot/internal/openai/enumset/jobstatus"
	"github.com/kylerqws/chatbot/internal/openai/enumset/model"
	"github.com/kylerqws/chatbot/internal/openai/enumset/purpose"
)

type OpenAiAdapter struct {
	command *cobra.Command
	enumset *enumset.ManagerSet
}

func NewOpenAiAdapter(cmd *cobra.Command) *OpenAiAdapter {
	return &OpenAiAdapter{command: cmd, enumset: enumset.NewManagerSet()}
}

func (h *OpenAiAdapter) FileStatusManager() *filestatus.Manager {
	return h.enumset.FileStatus()
}

func (h *OpenAiAdapter) ValidateFileStatusCode(code string) error {
	_, err := h.FileStatusManager().Resolve(code)
	if err != nil {
		return fmt.Errorf("invalid file status code '%s': %w", code, err)
	}
	return nil
}

func (h *OpenAiAdapter) JobStatusManager() *jobstatus.Manager {
	return h.enumset.JobStatus()
}

func (h *OpenAiAdapter) ValidateJobStatusCode(code string) error {
	_, err := h.JobStatusManager().Resolve(code)
	if err != nil {
		return fmt.Errorf("invalid job status code '%s': %w", code, err)
	}
	return nil
}

func (h *OpenAiAdapter) ModelManager() *model.Manager {
	return h.enumset.Model()
}

func (h *OpenAiAdapter) ValidateModelCode(code string) error {
	_, err := h.ModelManager().Resolve(code)
	if err != nil {
		return fmt.Errorf("invalid model code '%s': %w", code, err)
	}
	return nil
}

func (h *OpenAiAdapter) PurposeManager() *purpose.Manager {
	return h.enumset.Purpose()
}

func (h *OpenAiAdapter) ValidatePurposeCode(code string) error {
	_, err := h.PurposeManager().Resolve(code)
	if err != nil {
		return fmt.Errorf("invalid purpose code '%s': %w", code, err)
	}
	return nil
}
