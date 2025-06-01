package service

import "context"

type Job struct {
	ID              string   `json:"id"`
	Object          string   `json:"object"`
	CreatedAt       int64    `json:"created_at"`
	FinishedAt      int64    `json:"finished_at,omitempty"`
	Model           string   `json:"model"`
	FineTunedModel  string   `json:"fine_tuned_model,omitempty"`
	TrainingFile    string   `json:"training_file"`
	ValidationFile  string   `json:"validation_file,omitempty"`
	Status          string   `json:"status"`
	Hyperparameters any      `json:"hyperparameters,omitempty"`
	ResultFiles     []string `json:"result_files,omitempty"`
}

type CreateJobRequest struct {
	Model          string `json:"model"`
	TrainingFile   string `json:"training_file"`
	ValidationFile string `json:"validation_file,omitempty"`
}

type CreateJobResponse struct {
	Job *Job `json:"job"`
}

type GetJobInfoRequest struct {
	JobID string `json:"job_id"`
}

type GetJobInfoResponse struct {
	Job *Job `json:"job"`
}

type ListJobsRequest struct {
	JobIDs          []string `json:"job_ids,omitempty"`
	Models          []string `json:"models,omitempty"`
	FineTunedModels []string `json:"fine_tuned_models,omitempty"`
	Statuses        []string `json:"statuses,omitempty"`
	TrainingFiles   []string `json:"training_files,omitempty"`
	ValidationFiles []string `json:"validation_files,omitempty"`
	AfterJobID      string   `json:"after_job_id,omitempty"`
	LimitJobCount   uint8    `json:"limit_job_count,omitempty"`
	CreatedAfter    int64    `json:"created_after,omitempty"`
	CreatedBefore   int64    `json:"created_before,omitempty"`
	FinishedAfter   int64    `json:"finished_after,omitempty"`
	FinishedBefore  int64    `json:"finished_before,omitempty"`
}

type ListJobsResponse struct {
	Jobs    []*Job `json:"jobs"`
	HasMore bool   `json:"has_more"`
}

type CancelJobRequest struct {
	JobID string `json:"job_id"`
}

type CancelJobResponse struct {
	Job *Job `json:"job"`
}

type JobService interface {
	CreateJob(ctx context.Context, req *CreateJobRequest) (*CreateJobResponse, error)
	GetJobInfo(ctx context.Context, req *GetJobInfoRequest) (*GetJobInfoResponse, error)
	ListJobs(ctx context.Context, req *ListJobsRequest) (*ListJobsResponse, error)
	CancelJob(ctx context.Context, req *CancelJobRequest) (*CancelJobResponse, error)
}
