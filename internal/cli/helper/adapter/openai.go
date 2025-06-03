package adapter

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/spf13/cobra"

	"github.com/kylerqws/chatbot/internal/openai/enumset"
	"github.com/kylerqws/chatbot/internal/openai/enumset/filestatus"
	"github.com/kylerqws/chatbot/internal/openai/enumset/jobstatus"
	"github.com/kylerqws/chatbot/internal/openai/enumset/model"
	"github.com/kylerqws/chatbot/internal/openai/enumset/purpose"
)

const (
	AllFlagKey                   = "all"
	IdFlagKey                    = "id"
	StatusFlagKey                = "status"
	PurposeFlagKey               = "purpose"
	DefaultPurposeFlagKey        = "default-purpose"
	FilenameFlagKey              = "filename"
	ModelFlagKey                 = "model"
	DefaultModelFlagKey          = "default-model"
	FineTunedModelFlagKey        = "fine-tuned-model"
	TrainingFileFlagKey          = "training-file"
	ValidationFileFlagKey        = "validation-file"
	DefaultValidationFileFlagKey = "default-validation"
	CreatedAfterFlagKey          = "created-after"
	CreatedBeforeFlagKey         = "created-before"
	FinishedAfterFlagKey         = "finished-after"
	FinishedBeforeFlagKey        = "finished-before"
)

const (
	FileIDExample         = "file-xxxxxx..."
	JobIDExample          = "ftjob-xxxxxx..."
	PurposeExample        = "fine-tune-results"
	ModelExample          = "gpt-3.5-turbo-xxxx..."
	FineTunedModelExample = "ft:gpt-3.5-turbo-xxxx:personal::xxxxxx"
	Filename1Example      = "step_metrics.csv"
	Filename2Example      = "prompts.json"
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

func (*OpenAiAdapter) ValidateFileID(fileID string) error {
	if !regexp.MustCompile(`^file-[a-zA-Z0-9]{22,}$`).MatchString(fileID) {
		return fmt.Errorf("invalid file ID '%s'", fileID)
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

func (*OpenAiAdapter) ValidateJobID(jobID string) error {
	if !regexp.MustCompile(`^ftjob-[a-zA-Z0-9]{24,}$`).MatchString(jobID) {
		return fmt.Errorf("invalid job ID '%s'", jobID)
	}
	return nil
}

func (h *OpenAiAdapter) ModelManager() *model.Manager {
	return h.enumset.Model()
}

func (h *OpenAiAdapter) ValidateModelCode(code string) error {
	var bases []string

	for base := range h.ModelManager().List {
		bases = append(bases, regexp.QuoteMeta(base))
	}
	for i := range bases {
		if code == bases[i] || strings.HasPrefix(code, bases[i]+"-") {
			return nil
		}
	}

	pattern := fmt.Sprintf(`^ft:(%s)(?:-[a-zA-Z0-9]+)?:[a-z]+::[a-zA-Z0-9]+$`, strings.Join(bases, "|"))
	if matched, _ := regexp.MatchString(pattern, code); matched {
		return nil
	}

	return fmt.Errorf("invalid model code '%s'", code)
}

func (h *OpenAiAdapter) PurposeManager() *purpose.Manager {
	return h.enumset.Purpose()
}

func (h *OpenAiAdapter) ValidatePurposeCode(code string) error {
	var bases []string

	for base := range h.PurposeManager().List {
		bases = append(bases, regexp.QuoteMeta(base))
	}
	for i := range bases {
		if code == bases[i] || strings.HasPrefix(code, bases[i]+"-") {
			return nil
		}
	}

	pattern := fmt.Sprintf(`^(%s)(?:-[a-zA-Z0-9]+)*$`, strings.Join(bases, "|"))
	if matched, _ := regexp.MatchString(pattern, code); matched {
		return nil
	}

	return fmt.Errorf("invalid purpose code '%s'", code)
}
