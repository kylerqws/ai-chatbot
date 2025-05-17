package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/kylerqws/chatbot/pkg/openai/domain/model"
	"github.com/kylerqws/chatbot/pkg/openai/infrastructure/client"

	ctrcfg "github.com/kylerqws/chatbot/pkg/openai/contract/config"
	ctrsvc "github.com/kylerqws/chatbot/pkg/openai/contract/service"
)

type jobService struct {
	config ctrcfg.Config
	client *client.Client
}

func NewJobService(cl *client.Client, cfg ctrcfg.Config) ctrsvc.JobService {
	return &jobService{config: cfg, client: cl}
}

func (s *jobService) CreateJob(
	ctx context.Context,
	req *ctrsvc.CreateJobRequest,
) (*ctrsvc.CreateJobResponse, error) {
	result := &ctrsvc.CreateJobResponse{}

	mdl, err := model.Resolve(req.Model)
	if err != nil {
		return result, fmt.Errorf("failed to resolve model: %w", err)
	}
	req.Model = mdl.Code

	resp, err := s.client.RequestJSON(ctx, "POST", "/fine_tuning/jobs", req)
	if err != nil {
		return result, fmt.Errorf("failed to send request: %w", err)
	}

	var job ctrsvc.Job
	err = json.Unmarshal(resp, &job)
	if err != nil {
		return result, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	result.Job = &job
	return result, nil
}

func (s *jobService) GetJobInfo(
	ctx context.Context,
	req *ctrsvc.GetJobInfoRequest,
) (*ctrsvc.GetJobInfoResponse, error) {
	result := &ctrsvc.GetJobInfoResponse{}

	resp, err := s.client.Request(ctx, "GET", "/fine_tuning/jobs/"+req.JobID)
	if err != nil {
		return result, fmt.Errorf("failed to send request: %w", err)
	}

	var job ctrsvc.Job
	err = json.Unmarshal(resp, &job)
	if err != nil {
		return result, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	result.Job = &job
	return result, nil
}

func (s *jobService) ListJobs(
	ctx context.Context,
	req *ctrsvc.ListJobsRequest,
) (*ctrsvc.ListJobsResponse, error) {
	result := &ctrsvc.ListJobsResponse{}

	path := "/fine_tuning/jobs"
	if req.After != "" {
		params := url.Values{}
		params.Set("after", req.After)
		path += "?" + params.Encode()
	}

	resp, err := s.client.Request(ctx, "GET", path)
	if err != nil {
		return result, fmt.Errorf("failed to send request: %w", err)
	}

	var parsed struct {
		Data    []*ctrsvc.Job `json:"data"`
		HasMore bool          `json:"has_more"`
	}
	err = json.Unmarshal(resp, &parsed)
	if err != nil {
		return result, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	result.Jobs = s.filterJobs(parsed.Data, req)
	result.HasMore = parsed.HasMore

	return result, nil
}

func (s *jobService) CancelJob(
	ctx context.Context,
	req *ctrsvc.CancelJobRequest,
) (*ctrsvc.CancelJobResponse, error) {
	result := &ctrsvc.CancelJobResponse{}

	resp, err := s.client.Request(ctx, "POST", "/fine_tuning/jobs/"+req.JobID+"/cancel")
	if err != nil {
		return result, fmt.Errorf("failed to send request: %w", err)
	}

	var job ctrsvc.Job
	err = json.Unmarshal(resp, &job)
	if err != nil {
		return result, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	result.Job = &job
	return result, nil
}

func (_ *jobService) filterJobs(
	jobs []*ctrsvc.Job,
	req *ctrsvc.ListJobsRequest,
) []*ctrsvc.Job {
	var result []*ctrsvc.Job

	for _, j := range jobs {
		if (req.Model != "" && j.Model != req.Model) ||
			(req.Status != "" && j.Status != req.Status) ||
			(req.CreatedAfter > 0 && j.CreatedAt <= req.CreatedAfter) ||
			(req.CreatedBefore > 0 && j.CreatedAt >= req.CreatedBefore) {
			continue
		}

		result = append(result, j)
	}

	return result
}
