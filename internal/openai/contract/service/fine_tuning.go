package service

import (
	"context"
	ctrsvc "github.com/kylerqws/chatbot/pkg/openai/contract/service"
)

// FineTuningService defines operations for managing fine-tuning jobs in OpenAI.
type FineTuningService interface {
	// NewCreateJobRequest creates a new fine-tuning job creation request.
	NewCreateJobRequest() *ctrsvc.CreateJobRequest

	// NewCreateJobResponse creates a new fine-tuning job creation response.
	NewCreateJobResponse() *ctrsvc.CreateJobResponse

	// CreateJob creates a new fine-tuning job in OpenAI.
	CreateJob(ctx context.Context, req *ctrsvc.CreateJobRequest) (*ctrsvc.CreateJobResponse, error)

	// NewRetrieveJobRequest creates a new retrieve job request.
	NewRetrieveJobRequest() *ctrsvc.RetrieveJobRequest

	// NewRetrieveJobResponse creates a new retrieve job response.
	NewRetrieveJobResponse() *ctrsvc.RetrieveJobResponse

	// RetrieveJob retrieves a fine-tuning job metadata from OpenAI by ID.
	RetrieveJob(ctx context.Context, req *ctrsvc.RetrieveJobRequest) (*ctrsvc.RetrieveJobResponse, error)

	// NewPauseJobRequest creates a new pause job request.
	NewPauseJobRequest() *ctrsvc.PauseJobRequest

	// NewPauseJobResponse creates a new pause job response.
	NewPauseJobResponse() *ctrsvc.PauseJobResponse

	// PauseJob pauses a running fine-tuning job by ID.
	PauseJob(ctx context.Context, req *ctrsvc.PauseJobRequest) (*ctrsvc.PauseJobResponse, error)

	// NewResumeJobRequest creates a new resume job request.
	NewResumeJobRequest() *ctrsvc.ResumeJobRequest

	// NewResumeJobResponse creates a new resume job response.
	NewResumeJobResponse() *ctrsvc.ResumeJobResponse

	// ResumeJob resumes a paused fine-tuning job by ID.
	ResumeJob(ctx context.Context, req *ctrsvc.ResumeJobRequest) (*ctrsvc.ResumeJobResponse, error)

	// NewCancelJobRequest creates a new cancel job request.
	NewCancelJobRequest() *ctrsvc.CancelJobRequest

	// NewCancelJobResponse creates a new cancel job response.
	NewCancelJobResponse() *ctrsvc.CancelJobResponse

	// CancelJob cancels an active fine-tuning job by ID.
	CancelJob(ctx context.Context, req *ctrsvc.CancelJobRequest) (*ctrsvc.CancelJobResponse, error)

	// NewListJobsRequest creates a new list jobs request.
	NewListJobsRequest() *ctrsvc.ListJobsRequest

	// NewListJobsResponse creates a new list jobs response.
	NewListJobsResponse() *ctrsvc.ListJobsResponse

	// ListJobs retrieves a list of fine-tuning jobs from OpenAI.
	ListJobs(ctx context.Context, req *ctrsvc.ListJobsRequest) (*ctrsvc.ListJobsResponse, error)

	// NewListEventsRequest creates a new list events request.
	NewListEventsRequest() *ctrsvc.ListEventsRequest

	// NewListEventsResponse creates a new list events response.
	NewListEventsResponse() *ctrsvc.ListEventsResponse

	// ListEvents retrieves a list of events for a fine-tuning job.
	ListEvents(ctx context.Context, req *ctrsvc.ListEventsRequest) (*ctrsvc.ListEventsResponse, error)

	// NewListCheckpointsRequest creates a new list checkpoints request.
	NewListCheckpointsRequest() *ctrsvc.ListCheckpointsRequest

	// NewListCheckpointsResponse creates a new list checkpoints response.
	NewListCheckpointsResponse() *ctrsvc.ListCheckpointsResponse

	// ListCheckpoints retrieves a list of checkpoints for a fine-tuning job.
	ListCheckpoints(ctx context.Context, req *ctrsvc.ListCheckpointsRequest) (*ctrsvc.ListCheckpointsResponse, error)
}
