package adapter

import (
	"fmt"
	"github.com/spf13/cobra"

	"github.com/kylerqws/chatbot/internal/openai/enumset"
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

func (h *OpenAiAdapter) PurposeManager() *purpose.Manager {
	return h.enumset.Purpose()
}

func (h *OpenAiAdapter) ValidatePurposeCode(prpCode string) error {
	if _, err := h.enumset.Purpose().Resolve(prpCode); err != nil {
		return fmt.Errorf("invalid purpose code '%s': %w", prpCode, err)
	}
	return nil
}

func (h *OpenAiAdapter) ModelManager() *model.Manager {
	return h.enumset.Model()
}

func (h *OpenAiAdapter) ValidateModelCode(mdlCode string) error {
	if _, err := h.enumset.Model().Resolve(mdlCode); err != nil {
		return fmt.Errorf("invalid model code '%s': %w", mdlCode, err)
	}
	return nil
}

func (h *OpenAiAdapter) JobStatusManager() *jobstatus.Manager {
	return h.enumset.JobStatus()
}

func (h *OpenAiAdapter) ValidateJobStatusCode(stsCode string) error {
	if _, err := h.enumset.JobStatus().Resolve(stsCode); err != nil {
		return fmt.Errorf("invalid job status code '%s': %w", stsCode, err)
	}
	return nil
}
