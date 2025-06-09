package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/kylerqws/chatbot/pkg/openai/utils/query"

	ctrcl "github.com/kylerqws/chatbot/pkg/openai/contract/client"
	ctrcfg "github.com/kylerqws/chatbot/pkg/openai/contract/config"
	ctrsvc "github.com/kylerqws/chatbot/pkg/openai/contract/service"
)

// fineTuningService implements FineTuningService using OpenAI API client.
type fineTuningService struct {
	config ctrcfg.Config
	client ctrcl.Client
}

// NewFineTuningService creates a new FineTuningService instance.
func NewFineTuningService(cl ctrcl.Client, cfg ctrcfg.Config) ctrsvc.FineTuningService {
	return &fineTuningService{config: cfg, client: cl}
}

// CreateJob creates a new fine-tuning job on OpenAI.
func (s *fineTuningService) CreateJob(ctx context.Context, req *ctrsvc.CreateJobRequest) (*ctrsvc.CreateJobResponse, error) {
	result := &ctrsvc.CreateJobResponse{}

	resp, err := s.client.RequestJSON(ctx, "POST", "/fine_tuning/jobs", req)
	if err != nil {
		return result, fmt.Errorf("create job: %w", err)
	}

	var job ctrsvc.Job
	if err := json.Unmarshal(resp, &job); err != nil {
		return result, fmt.Errorf("unmarshal create job response: %w", err)
	}

	result.Job = &job
	return result, nil
}

// RetrieveJob fetches fine-tuning job metadata from OpenAI by its ID.
func (s *fineTuningService) RetrieveJob(ctx context.Context, req *ctrsvc.RetrieveJobRequest) (*ctrsvc.RetrieveJobResponse, error) {
	result := &ctrsvc.RetrieveJobResponse{}

	path := "/fine_tuning/jobs/" + url.PathEscape(req.JobID)
	resp, err := s.client.RequestRaw(ctx, "GET", path, nil)
	if err != nil {
		return result, fmt.Errorf("retrieve job: %w", err)
	}

	var job ctrsvc.Job
	if err := json.Unmarshal(resp, &job); err != nil {
		return result, fmt.Errorf("unmarshal retrieve job response: %w", err)
	}

	result.Job = &job
	return result, nil
}

// CancelJob sends a cancel request for a running fine-tuning job.
func (s *fineTuningService) CancelJob(ctx context.Context, req *ctrsvc.CancelJobRequest) (*ctrsvc.CancelJobResponse, error) {
	result := &ctrsvc.CancelJobResponse{}

	path := "/fine_tuning/jobs/" + url.PathEscape(req.JobID) + "/cancel"
	resp, err := s.client.RequestRaw(ctx, "POST", path, nil)
	if err != nil {
		return result, fmt.Errorf("cancel job: %w", err)
	}

	var job ctrsvc.Job
	if err := json.Unmarshal(resp, &job); err != nil {
		return result, fmt.Errorf("unmarshal cancel job response: %w", err)
	}

	result.Job = &job
	return result, nil
}

// PauseJob pauses an active fine-tuning job on OpenAI.
func (s *fineTuningService) PauseJob(ctx context.Context, req *ctrsvc.PauseJobRequest) (*ctrsvc.PauseJobResponse, error) {
	result := &ctrsvc.PauseJobResponse{}

	path := "/fine_tuning/jobs/" + url.PathEscape(req.JobID) + "/pause"
	resp, err := s.client.RequestRaw(ctx, "POST", path, nil)
	if err != nil {
		return result, fmt.Errorf("pause job: %w", err)
	}

	var job ctrsvc.Job
	if err := json.Unmarshal(resp, &job); err != nil {
		return result, fmt.Errorf("unmarshal pause job response: %w", err)
	}

	result.Job = &job
	return result, nil
}

// ResumeJob resumes a paused fine-tuning job on OpenAI.
func (s *fineTuningService) ResumeJob(ctx context.Context, req *ctrsvc.ResumeJobRequest) (*ctrsvc.ResumeJobResponse, error) {
	result := &ctrsvc.ResumeJobResponse{}

	path := "/fine_tuning/jobs/" + url.PathEscape(req.JobID) + "/resume"
	resp, err := s.client.RequestRaw(ctx, "POST", path, nil)
	if err != nil {
		return result, fmt.Errorf("resume job: %w", err)
	}

	var job ctrsvc.Job
	if err := json.Unmarshal(resp, &job); err != nil {
		return result, fmt.Errorf("unmarshal resume job response: %w", err)
	}

	result.Job = &job
	return result, nil
}

// ListJobs retrieves all fine-tuning jobs with optional pagination.
func (s *fineTuningService) ListJobs(ctx context.Context, req *ctrsvc.ListJobsRequest) (*ctrsvc.ListJobsResponse, error) {
	result := &ctrsvc.ListJobsResponse{}

	path := "/fine_tuning/jobs" + s.buildPaginationQuery(req.After, req.Limit)
	resp, err := s.client.RequestRaw(ctx, "GET", path, nil)
	if err != nil {
		return result, fmt.Errorf("list jobs: %w", err)
	}

	var parsed struct {
		Data []*ctrsvc.Job `json:"data"`
	}
	if err := json.Unmarshal(resp, &parsed); err != nil {
		return result, fmt.Errorf("unmarshal list jobs response: %w", err)
	}

	result.Jobs = parsed.Data
	return result, nil
}

// ListCheckpoints retrieves training checkpoints for a fine-tuning job.
func (s *fineTuningService) ListCheckpoints(ctx context.Context, req *ctrsvc.ListCheckpointsRequest) (*ctrsvc.ListCheckpointsResponse, error) {
	result := &ctrsvc.ListCheckpointsResponse{}

	path := "/fine_tuning/jobs/" + url.PathEscape(req.JobID) + "/checkpoints" + s.buildPaginationQuery(req.After, req.Limit)
	resp, err := s.client.RequestRaw(ctx, "GET", path, nil)
	if err != nil {
		return result, fmt.Errorf("list checkpoints: %w", err)
	}

	var parsed struct {
		Data []*ctrsvc.Checkpoint `json:"data"`
	}
	if err := json.Unmarshal(resp, &parsed); err != nil {
		return result, fmt.Errorf("unmarshal list checkpoints response: %w", err)
	}

	result.Checkpoints = parsed.Data
	return result, nil
}

// ListEvents retrieves event logs for a fine-tuning job.
func (s *fineTuningService) ListEvents(ctx context.Context, req *ctrsvc.ListEventsRequest) (*ctrsvc.ListEventsResponse, error) {
	result := &ctrsvc.ListEventsResponse{}

	path := "/fine_tuning/jobs/" + url.PathEscape(req.JobID) + "/events" + s.buildPaginationQuery(req.After, req.Limit)
	resp, err := s.client.RequestRaw(ctx, "GET", path, nil)
	if err != nil {
		return result, fmt.Errorf("list events: %w", err)
	}

	var parsed struct {
		Data []*ctrsvc.Event `json:"data"`
	}
	if err := json.Unmarshal(resp, &parsed); err != nil {
		return result, fmt.Errorf("unmarshal list events response: %w", err)
	}

	result.Events = parsed.Data
	return result, nil
}

// buildPaginationQuery constructs the API query string from the pagination parameters.
func (*fineTuningService) buildPaginationQuery(after *string, limit *uint8) string {
	q := query.NewUrlQuery()

	q.SetQueryStringParam("after", after)
	q.SetQueryUint8Param("limit", limit)

	return q.Encode()
}
