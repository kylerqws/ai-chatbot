package fine_tuning

import (
	"context"

	ctrint "github.com/kylerqws/chatbot/internal/openai/contract/service"
	ctrpkg "github.com/kylerqws/chatbot/pkg/openai/contract"
	ctrsvc "github.com/kylerqws/chatbot/pkg/openai/contract/service"
)

// service provides operations for managing fine-tuning jobs in OpenAI.
type service struct {
	ctx context.Context
	svc ctrsvc.FineTuningService
}

// NewService creates a new fine-tuning service for managing OpenAI fine-tuning jobs.
func NewService(ctx context.Context, sdk ctrpkg.OpenAI) ctrint.FineTuningService {
	return &service{ctx: ctx, svc: sdk.FineTuningService()}
}

// NewCreateJobRequest creates a new fine-tuning job creation request.
func (*service) NewCreateJobRequest() *ctrsvc.CreateJobRequest {
	return &ctrsvc.CreateJobRequest{}
}

// NewCreateJobResponse creates a new fine-tuning job creation response.
func (*service) NewCreateJobResponse() *ctrsvc.CreateJobResponse {
	return &ctrsvc.CreateJobResponse{}
}

// CreateJob creates a new fine-tuning job in OpenAI.
func (s *service) CreateJob(ctx context.Context, req *ctrsvc.CreateJobRequest) (*ctrsvc.CreateJobResponse, error) {
	return s.svc.CreateJob(ctx, req)
}

// NewRetrieveJobRequest creates a new retrieve job request.
func (*service) NewRetrieveJobRequest() *ctrsvc.RetrieveJobRequest {
	return &ctrsvc.RetrieveJobRequest{}
}

// NewRetrieveJobResponse creates a new retrieve job response.
func (*service) NewRetrieveJobResponse() *ctrsvc.RetrieveJobResponse {
	return &ctrsvc.RetrieveJobResponse{}
}

// RetrieveJob retrieves a fine-tuning job metadata from OpenAI by ID.
func (s *service) RetrieveJob(ctx context.Context, req *ctrsvc.RetrieveJobRequest) (*ctrsvc.RetrieveJobResponse, error) {
	return s.svc.RetrieveJob(ctx, req)
}

// NewPauseJobRequest creates a new pause job request.
func (*service) NewPauseJobRequest() *ctrsvc.PauseJobRequest {
	return &ctrsvc.PauseJobRequest{}
}

// NewPauseJobResponse creates a new pause job response.
func (*service) NewPauseJobResponse() *ctrsvc.PauseJobResponse {
	return &ctrsvc.PauseJobResponse{}
}

// PauseJob pauses a running fine-tuning job by ID.
func (s *service) PauseJob(ctx context.Context, req *ctrsvc.PauseJobRequest) (*ctrsvc.PauseJobResponse, error) {
	return s.svc.PauseJob(ctx, req)
}

// NewResumeJobRequest creates a new resume job request.
func (*service) NewResumeJobRequest() *ctrsvc.ResumeJobRequest {
	return &ctrsvc.ResumeJobRequest{}
}

// NewResumeJobResponse creates a new resume job response.
func (*service) NewResumeJobResponse() *ctrsvc.ResumeJobResponse {
	return &ctrsvc.ResumeJobResponse{}
}

// ResumeJob resumes a paused fine-tuning job by ID.
func (s *service) ResumeJob(ctx context.Context, req *ctrsvc.ResumeJobRequest) (*ctrsvc.ResumeJobResponse, error) {
	return s.svc.ResumeJob(ctx, req)
}

// NewCancelJobRequest creates a new cancel job request.
func (*service) NewCancelJobRequest() *ctrsvc.CancelJobRequest {
	return &ctrsvc.CancelJobRequest{}
}

// NewCancelJobResponse creates a new cancel job response.
func (*service) NewCancelJobResponse() *ctrsvc.CancelJobResponse {
	return &ctrsvc.CancelJobResponse{}
}

// CancelJob cancels an active fine-tuning job by ID.
func (s *service) CancelJob(ctx context.Context, req *ctrsvc.CancelJobRequest) (*ctrsvc.CancelJobResponse, error) {
	return s.svc.CancelJob(ctx, req)
}

// NewListJobsRequest creates a new list jobs request.
func (*service) NewListJobsRequest() *ctrsvc.ListJobsRequest {
	return &ctrsvc.ListJobsRequest{}
}

// NewListJobsResponse creates a new list jobs response.
func (*service) NewListJobsResponse() *ctrsvc.ListJobsResponse {
	return &ctrsvc.ListJobsResponse{}
}

// ListJobs retrieves a list of fine-tuning jobs from OpenAI.
func (s *service) ListJobs(ctx context.Context, req *ctrsvc.ListJobsRequest) (*ctrsvc.ListJobsResponse, error) {
	return s.svc.ListJobs(ctx, req)
}

// NewListEventsRequest creates a new list events request.
func (*service) NewListEventsRequest() *ctrsvc.ListEventsRequest {
	return &ctrsvc.ListEventsRequest{}
}

// NewListEventsResponse creates a new list events response.
func (*service) NewListEventsResponse() *ctrsvc.ListEventsResponse {
	return &ctrsvc.ListEventsResponse{}
}

// ListEvents retrieves a list of events for a fine-tuning job.
func (s *service) ListEvents(ctx context.Context, req *ctrsvc.ListEventsRequest) (*ctrsvc.ListEventsResponse, error) {
	return s.svc.ListEvents(ctx, req)
}

// NewListCheckpointsRequest creates a new list checkpoints request.
func (*service) NewListCheckpointsRequest() *ctrsvc.ListCheckpointsRequest {
	return &ctrsvc.ListCheckpointsRequest{}
}

// NewListCheckpointsResponse creates a new list checkpoints response.
func (*service) NewListCheckpointsResponse() *ctrsvc.ListCheckpointsResponse {
	return &ctrsvc.ListCheckpointsResponse{}
}

// ListCheckpoints retrieves a list of checkpoints for a fine-tuning job.
func (s *service) ListCheckpoints(ctx context.Context, req *ctrsvc.ListCheckpointsRequest) (*ctrsvc.ListCheckpointsResponse, error) {
	return s.svc.ListCheckpoints(ctx, req)
}
