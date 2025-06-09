package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/kylerqws/chatbot/pkg/openai/utils/query"

	ctrcli "github.com/kylerqws/chatbot/pkg/openai/contract/client"
	ctrcfg "github.com/kylerqws/chatbot/pkg/openai/contract/config"
	ctrsvc "github.com/kylerqws/chatbot/pkg/openai/contract/service"
)

// fineTuningService implements FineTuningService using the OpenAI API client.
type fineTuningService struct {
	config ctrcfg.Config
	client ctrcli.Client
}

// NewFineTuningService creates a new instance of FineTuningService.
func NewFineTuningService(cl ctrcli.Client, cfg ctrcfg.Config) ctrsvc.FineTuningService {
	return &fineTuningService{config: cfg, client: cl}
}

// CreateJob starts a new fine-tuning job.
func (s *fineTuningService) CreateJob(ctx context.Context, req *ctrsvc.CreateFineTuningJobRequest) (*ctrsvc.CreateFineTuningJobResponse, error) {
	result := &ctrsvc.CreateFineTuningJobResponse{}

	resp, err := s.client.RequestJSON(ctx, "POST", "/fine_tuning/jobs", req)
	if err != nil {
		return nil, fmt.Errorf("create job: %w", err)
	}

	if err := json.Unmarshal(resp, result); err != nil {
		return nil, fmt.Errorf("unmarshal create job response: %w", err)
	}

	return result, nil
}

// RetrieveJob fetches details of a fine-tuning job by ID.
func (s *fineTuningService) RetrieveJob(ctx context.Context, req *ctrsvc.RetrieveFineTuningJobRequest) (*ctrsvc.RetrieveFineTuningJobResponse, error) {
	result := &ctrsvc.RetrieveFineTuningJobResponse{}

	path := "/fine_tuning/jobs/" + url.PathEscape(req.JobID)
	resp, err := s.client.RequestRaw(ctx, "GET", path, nil)
	if err != nil {
		return nil, fmt.Errorf("retrieve job: %w", err)
	}

	if err := json.Unmarshal(resp, result); err != nil {
		return nil, fmt.Errorf("unmarshal retrieve job response: %w", err)
	}

	return result, nil
}

// CancelJob cancels a running fine-tuning job.
func (s *fineTuningService) CancelJob(ctx context.Context, req *ctrsvc.CancelFineTuningJobRequest) (*ctrsvc.CancelFineTuningJobResponse, error) {
	result := &ctrsvc.CancelFineTuningJobResponse{}

	path := "/fine_tuning/jobs/" + url.PathEscape(req.JobID) + "/cancel"
	resp, err := s.client.RequestRaw(ctx, "POST", path, nil)
	if err != nil {
		return nil, fmt.Errorf("cancel job: %w", err)
	}

	if err := json.Unmarshal(resp, result); err != nil {
		return nil, fmt.Errorf("unmarshal cancel job response: %w", err)
	}

	return result, nil
}

// PauseJob pauses a running fine-tuning job.
func (s *fineTuningService) PauseJob(ctx context.Context, req *ctrsvc.PauseFineTuningJobRequest) (*ctrsvc.PauseFineTuningJobResponse, error) {
	result := &ctrsvc.PauseFineTuningJobResponse{}

	path := "/fine_tuning/jobs/" + url.PathEscape(req.JobID) + "/pause"
	resp, err := s.client.RequestRaw(ctx, "POST", path, nil)
	if err != nil {
		return nil, fmt.Errorf("pause job: %w", err)
	}

	if err := json.Unmarshal(resp, result); err != nil {
		return nil, fmt.Errorf("unmarshal pause job response: %w", err)
	}

	return result, nil
}

// ResumeJob resumes a paused fine-tuning job.
func (s *fineTuningService) ResumeJob(ctx context.Context, req *ctrsvc.ResumeFineTuningJobRequest) (*ctrsvc.ResumeFineTuningJobResponse, error) {
	result := &ctrsvc.ResumeFineTuningJobResponse{}

	path := "/fine_tuning/jobs/" + url.PathEscape(req.JobID) + "/resume"
	resp, err := s.client.RequestRaw(ctx, "POST", path, nil)
	if err != nil {
		return nil, fmt.Errorf("resume job: %w", err)
	}

	if err := json.Unmarshal(resp, result); err != nil {
		return nil, fmt.Errorf("unmarshal resume job response: %w", err)
	}

	return result, nil
}

// ListJobs returns all fine-tuning jobs with optional pagination.
func (s *fineTuningService) ListJobs(ctx context.Context, req *ctrsvc.ListFineTuningJobsRequest) (*ctrsvc.ListFineTuningJobsResponse, error) {
	result := &ctrsvc.ListFineTuningJobsResponse{}

	path := "/fine_tuning/jobs" + s.buildQuery(req.After, req.Limit)
	resp, err := s.client.RequestRaw(ctx, "GET", path, nil)
	if err != nil {
		return nil, fmt.Errorf("list jobs: %w", err)
	}

	var parsed struct {
		Data []*ctrsvc.FineTuningJob `json:"data"`
	}
	if err := json.Unmarshal(resp, &parsed); err != nil {
		return nil, fmt.Errorf("unmarshal list jobs response: %w", err)
	}

	result.Jobs = parsed.Data
	return result, nil
}

// ListCheckpoints returns training checkpoints for a fine-tuning job.
func (s *fineTuningService) ListCheckpoints(ctx context.Context, req *ctrsvc.ListFineTuningJobCheckpointsRequest) (*ctrsvc.ListFineTuningJobCheckpointsResponse, error) {
	result := &ctrsvc.ListFineTuningJobCheckpointsResponse{}

	path := "/fine_tuning/jobs/" + url.PathEscape(req.FineTuningJobID) + "/checkpoints" + s.buildQuery(req.After, req.Limit)
	resp, err := s.client.RequestRaw(ctx, "GET", path, nil)
	if err != nil {
		return nil, fmt.Errorf("list checkpoints: %w", err)
	}

	if err := json.Unmarshal(resp, result); err != nil {
		return nil, fmt.Errorf("unmarshal list checkpoints response: %w", err)
	}

	return result, nil
}

// ListEvents returns event logs for a fine-tuning job.
func (s *fineTuningService) ListEvents(ctx context.Context, req *ctrsvc.ListFineTuningJobEventsRequest) (*ctrsvc.ListFineTuningJobEventsResponse, error) {
	result := &ctrsvc.ListFineTuningJobEventsResponse{}

	path := "/fine_tuning/jobs/" + url.PathEscape(req.FineTuningJobID) + "/events" + s.buildQuery(req.After, req.Limit)
	resp, err := s.client.RequestRaw(ctx, "GET", path, nil)
	if err != nil {
		return nil, fmt.Errorf("list events: %w", err)
	}

	if err := json.Unmarshal(resp, result); err != nil {
		return nil, fmt.Errorf("unmarshal list events response: %w", err)
	}

	return result, nil
}

func (s *fineTuningService) buildQuery(after *string, limit *uint8) string {
	q := query.NewUrlQuery()

	q.SetQueryStringParam("after", after)
	q.SetQueryUint8Param("after", limit)

	return q.Encode()
}
