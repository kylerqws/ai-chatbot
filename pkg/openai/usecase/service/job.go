package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/kylerqws/chatbot/pkg/openai/domain/model"
	"github.com/kylerqws/chatbot/pkg/openai/infrastructure/client"
	"github.com/kylerqws/chatbot/pkg/openai/utils/filter"

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

	path := "/fine_tuning/jobs" + s.buildListJobsQuery(req)
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

	result.Jobs = s.applyListJobsFilter(parsed.Data, req)
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

func (*jobService) buildListJobsQuery(req *ctrsvc.ListJobsRequest) string {
	params := url.Values{}

	if req.AfterJobID != "" {
		params.Set("after", req.AfterJobID)
	}
	if req.LimitJobs != 0 {
		params.Set("limit", strconv.FormatUint(uint64(req.LimitJobs), 10))
	}

	if query := params.Encode(); query != "" {
		return "?" + query
	}
	return ""
}

func (*jobService) hasAnyListJobsFilter(req *ctrsvc.ListJobsRequest) bool {
	return req.CreatedAfter != 0 || req.CreatedBefore != 0 ||
		req.FinishedAfter != 0 || req.FinishedBefore != 0 ||
		len(req.JobIDs) > 0 || len(req.Statuses) > 0 ||
		len(req.Models) > 0 || len(req.FineTunedModels) > 0 ||
		len(req.TrainingFiles) > 0 || len(req.ValidationFiles) > 0
}

func (s *jobService) applyListJobsFilter(jobs []*ctrsvc.Job, req *ctrsvc.ListJobsRequest) []*ctrsvc.Job {
	if !s.hasAnyListJobsFilter(req) {
		return jobs
	}

	var result []*ctrsvc.Job
	for i := range jobs {
		if !filter.MatchDateValue(jobs[i].CreatedAt, req.CreatedAfter, req.CreatedBefore) {
			continue
		}
		if !filter.MatchDateValue(jobs[i].FinishedAt, req.FinishedAfter, req.FinishedBefore) {
			continue
		}
		if !filter.MatchStrValue(jobs[i].ID, req.JobIDs) {
			continue
		}
		if !filter.MatchStrValue(jobs[i].Status, req.Statuses) {
			continue
		}
		if !filter.MatchStrValue(jobs[i].Model, req.Models) {
			continue
		}
		if !filter.MatchStrValue(jobs[i].FineTunedModel, req.FineTunedModels) {
			continue
		}
		if !filter.MatchStrValue(jobs[i].TrainingFile, req.TrainingFiles) {
			continue
		}
		if !filter.MatchStrValue(jobs[i].ValidationFile, req.ValidationFiles) {
			continue
		}

		result = append(result, jobs[i])
	}

	return result
}
