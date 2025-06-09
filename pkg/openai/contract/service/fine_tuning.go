package service

import (
	"context"
	"github.com/kylerqws/chatbot/pkg/openai/utils/value"
)

// FineTuningJob represents a fine-tuning job resource.
type FineTuningJob struct {
	ID                string                  `json:"id"`
	Object            string                  `json:"object"`
	Status            string                  `json:"status"`
	Model             string                  `json:"model"`
	TrainingFile      string                  `json:"training_file"`
	CreatedAt         int64                   `json:"created_at"`
	UpdatedAt         int64                   `json:"updated_at"`
	OrganizationID    string                  `json:"organization_id"`
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

// FineTuningPermission describes permissions for a fine-tuning job.
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

// FineTuningJobEvent represents a log event for a fine-tuning job.
type FineTuningJobEvent struct {
	ID        string `json:"id"`
	Object    string `json:"object"`
	Level     string `json:"level"`
	Message   string `json:"message"`
	CreatedAt int64  `json:"created_at"`
}

// FineTuningJobCheckpoint represents a checkpoint of a fine-tuning job.
type FineTuningJobCheckpoint struct {
	ID                       string                 `json:"id"`
	Object                   string                 `json:"object"`
	FineTuningJobID          string                 `json:"fine_tuning_job_id"`
	Metrics                  map[string]interface{} `json:"metrics"`
	EvalMetrics              map[string]interface{} `json:"eval_metrics"`
	StepNumber               int64                  `json:"step_number"`
	CreatedAt                int64                  `json:"created_at"`
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

// CreateFineTuningJobRequest contains parameters to create a fine-tuning job.
type CreateFineTuningJobRequest struct {
	TrainingFile       string            `json:"training_file"`
	Model              *string           `json:"model,omitempty"`
	Suffix             *string           `json:"suffix,omitempty"`
	ValidationFile     *string           `json:"validation_file,omitempty"`
	IntegrationType    *string           `json:"integration_type,omitempty"`
	IntegrationMapping map[string]string `json:"integration_mapping,omitempty"`
	Hyperparameters    *Hyperparameters  `json:"hyperparameters,omitempty"`
}

// CreateFineTuningJobResponse wraps the created fine-tuning job returned from the API.
type CreateFineTuningJobResponse struct {
	FineTuningJob *FineTuningJob `json:"fine_tuning_job"`
}

// RetrieveFineTuningJobRequest contains the ID of the fine-tuning job to retrieve.
type RetrieveFineTuningJobRequest struct {
	FineTuningJobID string `json:"fine_tuning_job_id"`
}

// RetrieveFineTuningJobResponse wraps the fine-tuning job metadata returned from the API.
type RetrieveFineTuningJobResponse struct {
	FineTuningJob *FineTuningJob `json:"fine_tuning_job"`
}

// CancelFineTuningJobRequest contains the ID of the fine-tuning job to cancel.
type CancelFineTuningJobRequest struct {
	FineTuningJobID string `json:"fine_tuning_job_id"`
}

// CancelFineTuningJobResponse wraps the canceled fine-tuning job returned from the API.
type CancelFineTuningJobResponse struct {
	FineTuningJob *FineTuningJob `json:"fine_tuning_job"`
}

// ListFineTuningJobsRequest contains parameters for filtering listed fine-tuning jobs.
type ListFineTuningJobsRequest struct {
	// API-supported query parameters
	After *string `json:"after,omitempty"`
	Limit *uint8  `json:"limit,omitempty"`
}

// ListFineTuningJobsResponse wraps a list of fine-tuning jobs returned from the API.
type ListFineTuningJobsResponse struct {
	FineTuningJobs []*FineTuningJob `json:"fine_tuning_jobs"`
}

// ListFineTuningJobEventsRequest contains parameters for filtering listed job events.
type ListFineTuningJobEventsRequest struct {
	// API-supported query parameters
	FineTuningJobID string  `json:"fine_tuning_job_id"`
	After           *string `json:"after,omitempty"`
	Limit           *uint8  `json:"limit,omitempty"`
}

// ListFineTuningJobEventsResponse wraps a list of job events returned from the API.
type ListFineTuningJobEventsResponse struct {
	Events []*FineTuningJobEvent `json:"events"`
}

// ListFineTuningJobCheckpointsRequest contains parameters for filtering listed job checkpoints.
type ListFineTuningJobCheckpointsRequest struct {
	// API-supported query parameters
	FineTuningJobID string  `json:"fine_tuning_job_id"`
	After           *string `json:"after,omitempty"`
	Limit           *uint8  `json:"limit,omitempty"`
}

// ListFineTuningJobCheckpointsResponse wraps a list of job checkpoints returned from the API.
type ListFineTuningJobCheckpointsResponse struct {
	Checkpoints []*FineTuningJobCheckpoint `json:"checkpoints"`
}

// PauseFineTuningJobRequest contains the ID of the fine-tuning job to pause.
type PauseFineTuningJobRequest struct {
	FineTuningJobID string `json:"fine_tuning_job_id"`
}

// PauseFineTuningJobResponse wraps the paused fine-tuning job returned from the API.
type PauseFineTuningJobResponse struct {
	FineTuningJob *FineTuningJob `json:"fine_tuning_job"`
}

// ResumeFineTuningJobRequest contains the ID of the fine-tuning job to resume.
type ResumeFineTuningJobRequest struct {
	FineTuningJobID string `json:"fine_tuning_job_id"`
}

// ResumeFineTuningJobResponse wraps the resumed fine-tuning job returned from the API.
type ResumeFineTuningJobResponse struct {
	FineTuningJob *FineTuningJob `json:"fine_tuning_job"`
}

// FineTuningService defines operations for managing fine-tuning jobs.
type FineTuningService interface {
	// CreateJob creates a new fine-tuning job.
	CreateJob(ctx context.Context, req *CreateFineTuningJobRequest) (*CreateFineTuningJobResponse, error)

	// RetrieveJob retrieves a fine-tuning job by its ID.
	RetrieveJob(ctx context.Context, req *RetrieveFineTuningJobRequest) (*RetrieveFineTuningJobResponse, error)

	// CancelJob cancels an active fine-tuning job.
	CancelJob(ctx context.Context, req *CancelFineTuningJobRequest) (*CancelFineTuningJobResponse, error)

	// PauseJob pauses a running fine-tuning job.
	PauseJob(ctx context.Context, req *PauseFineTuningJobRequest) (*PauseFineTuningJobResponse, error)

	// ResumeJob resumes a paused fine-tuning job.
	ResumeJob(ctx context.Context, req *ResumeFineTuningJobRequest) (*ResumeFineTuningJobResponse, error)

	// ListJobs returns a list of fine-tuning jobs with optional pagination.
	ListJobs(ctx context.Context, req *ListFineTuningJobsRequest) (*ListFineTuningJobsResponse, error)

	// ListCheckpoints returns checkpoints for a specific fine-tuning job.
	ListCheckpoints(ctx context.Context, req *ListFineTuningJobCheckpointsRequest) (*ListFineTuningJobCheckpointsResponse, error)

	// ListEvents returns events for a specific fine-tuning job.
	ListEvents(ctx context.Context, req *ListFineTuningJobEventsRequest) (*ListFineTuningJobEventsResponse, error)
}
