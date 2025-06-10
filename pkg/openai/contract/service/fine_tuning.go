package service

import (
	"context"
	"github.com/kylerqws/chatbot/pkg/openai/utils/value"
)

// Job represents a job object returned from the OpenAI API.
type Job struct {
	ID                string                  `json:"id"`
	Object            string                  `json:"object"`
	Model             string                  `json:"model"`
	TrainingFile      string                  `json:"training_file"`
	OrganizationID    string                  `json:"organization_id"`
	Status            string                  `json:"status"`
	CreatedAt         int64                   `json:"created_at"`
	UpdatedAt         int64                   `json:"updated_at"`
	Hyperparameters   map[string]interface{}  `json:"hyperparameters,omitempty"`
	ResultFiles       []string                `json:"result_files"`
	FineTunedModel    *string                 `json:"fine_tuned_model,omitempty"`
	ValidationFile    *string                 `json:"validation_file,omitempty"`
	FailureReason     *string                 `json:"failure_reason,omitempty"`
	Error             *string                 `json:"error,omitempty"`
	Integration       *string                 `json:"integration,omitempty"`
	Suffix            *string                 `json:"suffix,omitempty"`
	FinishedAt        *int64                  `json:"finished_at,omitempty"`
	EstimatedFinishAt *int64                  `json:"estimated_finish,omitempty"`
	TrainedTokens     *int64                  `json:"trained_tokens,omitempty"`
	Seed              *int                    `json:"seed,omitempty"`
	Permissions       []*FineTuningPermission `json:"permissions,omitempty"`
}

// FineTuningPermission represents a permission object returned from the OpenAI API.
type FineTuningPermission struct {
	ID                 string  `json:"id"`
	Object             string  `json:"object"`
	Organization       string  `json:"organization"`
	IsBlocking         bool    `json:"is_blocking"`
	AllowCreateEngine  bool    `json:"allow_create_engine"`
	AllowSampling      bool    `json:"allow_sampling"`
	AllowLogprobs      bool    `json:"allow_logprobs"`
	AllowSearchIndices bool    `json:"allow_search_indices"`
	AllowView          bool    `json:"allow_view"`
	AllowFineTuning    bool    `json:"allow_fine_tuning"`
	CreatedAt          int64   `json:"created_at"`
	Group              *string `json:"group,omitempty"`
}

// Event represents an event object returned from the OpenAI API.
type Event struct {
	ID        string `json:"id"`
	Object    string `json:"object"`
	Level     string `json:"level"`
	Message   string `json:"message"`
	CreatedAt int64  `json:"created_at"`
}

// Checkpoint represents a checkpoint object returned from the OpenAI API.
type Checkpoint struct {
	ID                       string                 `json:"id"`
	Object                   string                 `json:"object"`
	StepNumber               int64                  `json:"step_number"`
	CreatedAt                int64                  `json:"created_at"`
	Metrics                  map[string]interface{} `json:"metrics"`
	EvalMetrics              map[string]interface{} `json:"eval_metrics"`
	FullModelFile            *string                `json:"full_model_file,omitempty"`
	FineTunedModelCheckpoint *string                `json:"fine_tuned_model_checkpoint,omitempty"`
}

// FineTuningChatInput is a training input of type chat.
type FineTuningChatInput struct {
	Messages []*FineTuningChatMessage `json:"messages"`
}

// FineTuningChatMessage represents a single chat message with role and content.
type FineTuningChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// FineTuningPreferenceInput is a training example for preference-based fine-tuning.
type FineTuningPreferenceInput struct {
	Prompt     string `json:"prompt"`
	Completion string `json:"completion"`
	Rating     int    `json:"rating"`
}

// FineTuningReinforcementInput is a training example for reinforcement learning-based fine-tuning.
type FineTuningReinforcementInput struct {
	Prompt     string `json:"prompt"`
	Completion string `json:"completion"`
	Reward     int    `json:"reward"`
}

// Hyperparameters holds fine-tuning hyperparameters.
type Hyperparameters struct {
	NEpochs                *value.AutoOrNumber[int]     `json:"n_epochs,omitempty"`
	BatchSize              *value.AutoOrNumber[int]     `json:"batch_size,omitempty"`
	LearningRateMultiplier *value.AutoOrNumber[float64] `json:"learning_rate_multiplier,omitempty"`
}

// CreateJobRequest contains parameters to create a fine-tuning job.
type CreateJobRequest struct {
	TrainingFile       string            `json:"training_file"`
	Model              *string           `json:"model,omitempty"`
	Suffix             *string           `json:"suffix,omitempty"`
	ValidationFile     *string           `json:"validation_file,omitempty"`
	IntegrationType    *string           `json:"integration_type,omitempty"`
	IntegrationMapping map[string]string `json:"integration_mapping,omitempty"`
	Hyperparameters    *Hyperparameters  `json:"hyperparameters,omitempty"`
}

// CreateJobResponse wraps the created fine-tuning job returned from the API.
type CreateJobResponse struct {
	Job *Job `json:"job"`
}

// RetrieveJobRequest contains the ID of the fine-tuning job to retrieve.
type RetrieveJobRequest struct {
	JobID string `json:"job_id"`
}

// RetrieveJobResponse wraps the fine-tuning job metadata returned from the API.
type RetrieveJobResponse struct {
	Job *Job `json:"job"`
}

// CancelJobRequest contains the ID of the fine-tuning job to cancel.
type CancelJobRequest struct {
	JobID string `json:"job_id"`
}

// CancelJobResponse wraps the canceled fine-tuning job returned from the API.
type CancelJobResponse struct {
	Job *Job `json:"job"`
}

// ListJobsRequest contains parameters for filtering listed fine-tuning jobs.
type ListJobsRequest struct {
	// Local filtering (applied after fetching data)
	JobIDs                []string `json:"job_ids,omitempty"`
	OrganizationIDs       []string `json:"organization_ids,omitempty"`
	Statuses              []string `json:"statuses,omitempty"`
	Suffixes              []string `json:"suffixes,omitempty"`
	Models                []string `json:"models,omitempty"`
	FineTunedModels       []string `json:"fine_tuned_models,omitempty"`
	TrainingFiles         []string `json:"training_files,omitempty"`
	ValidationFiles       []string `json:"validation_files,omitempty"`
	CreatedAfter          *int64   `json:"created_after,omitempty"`
	CreatedBefore         *int64   `json:"created_before,omitempty"`
	UpdatedAfter          *int64   `json:"updated_after,omitempty"`
	UpdatedBefore         *int64   `json:"updated_before,omitempty"`
	FinishedAfter         *int64   `json:"finished_after,omitempty"`
	FinishedBefore        *int64   `json:"finished_before,omitempty"`
	EstimatedFinishAfter  *int64   `json:"estimated_finish_after,omitempty"`
	EstimatedFinishBefore *int64   `json:"estimated_finish_before,omitempty"`

	// API-supported query parameters
	After *string `json:"after,omitempty"`
	Limit *uint8  `json:"limit,omitempty"`
}

// ListJobsResponse wraps a list of fine-tuning jobs returned from the API.
type ListJobsResponse struct {
	Jobs []*Job `json:"jobs"`
}

// ListEventsRequest contains parameters for filtering listed fine-tuning job events.
type ListEventsRequest struct {
	// Local filtering (applied after fetching data)
	EventIDs      []string `json:"event_ids,omitempty"`
	Levels        []string `json:"levels,omitempty"`
	CreatedAfter  *int64   `json:"created_after,omitempty"`
	CreatedBefore *int64   `json:"created_before,omitempty"`

	// API-supported query parameters
	JobID string  `json:"job_id"`
	After *string `json:"after,omitempty"`
	Limit *uint8  `json:"limit,omitempty"`
}

// ListEventsResponse wraps a list of fine-tuning job events returned from the API.
type ListEventsResponse struct {
	Events []*Event `json:"events"`
}

// ListCheckpointsRequest contains parameters for filtering listed fine-tuning job checkpoints.
type ListCheckpointsRequest struct {
	// Local filtering (applied after fetching data)
	CheckpointIDs []string `json:"checkpoint_ids,omitempty"`
	CreatedAfter  *int64   `json:"created_after,omitempty"`
	CreatedBefore *int64   `json:"created_before,omitempty"`

	// API-supported query parameters
	JobID string  `json:"job_id"`
	After *string `json:"after,omitempty"`
	Limit *uint8  `json:"limit,omitempty"`
}

// ListCheckpointsResponse wraps a list of fine-tuning job checkpoints returned from the API.
type ListCheckpointsResponse struct {
	Checkpoints []*Checkpoint `json:"checkpoints"`
}

// PauseJobRequest contains the ID of the fine-tuning job to pause.
type PauseJobRequest struct {
	JobID string `json:"job_id"`
}

// PauseJobResponse wraps the paused fine-tuning job returned from the API.
type PauseJobResponse struct {
	Job *Job `json:"job"`
}

// ResumeJobRequest contains the ID of the fine-tuning job to resume.
type ResumeJobRequest struct {
	JobID string `json:"job_id"`
}

// ResumeJobResponse wraps the resumed fine-tuning job returned from the API.
type ResumeJobResponse struct {
	Job *Job `json:"job"`
}

// FineTuningService defines operations for managing fine-tuning jobs in OpenAI.
type FineTuningService interface {
	// CreateJob creates a new fine-tuning job in OpenAI.
	CreateJob(ctx context.Context, req *CreateJobRequest) (*CreateJobResponse, error)

	// RetrieveJob retrieves a fine-tuning job from OpenAI by its ID.
	RetrieveJob(ctx context.Context, req *RetrieveJobRequest) (*RetrieveJobResponse, error)

	// CancelJob cancels an active fine-tuning job in OpenAI.
	CancelJob(ctx context.Context, req *CancelJobRequest) (*CancelJobResponse, error)

	// PauseJob pauses a running fine-tuning job in OpenAI.
	PauseJob(ctx context.Context, req *PauseJobRequest) (*PauseJobResponse, error)

	// ResumeJob resumes a paused fine-tuning job in OpenAI.
	ResumeJob(ctx context.Context, req *ResumeJobRequest) (*ResumeJobResponse, error)

	// ListJobs returns a filtered list of fine-tuning jobs from OpenAI.
	ListJobs(ctx context.Context, req *ListJobsRequest) (*ListJobsResponse, error)

	// ListEvents returns a filtered list of events for a fine-tuning job from OpenAI.
	ListEvents(ctx context.Context, req *ListEventsRequest) (*ListEventsResponse, error)

	// ListCheckpoints returns a filtered list of checkpoints for a fine-tuning job from OpenAI.
	ListCheckpoints(ctx context.Context, req *ListCheckpointsRequest) (*ListCheckpointsResponse, error)
}
