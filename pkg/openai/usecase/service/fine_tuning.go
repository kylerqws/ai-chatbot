package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/kylerqws/chatbot/pkg/openai/utils/filter"
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

// RetrieveJob retrieves fine-tuning job metadata from OpenAI by its ID.
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

// PauseJob pauses a running fine-tuning job in OpenAI by its ID.
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

// ResumeJob resumes a paused fine-tuning job in OpenAI by its ID.
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

// CancelJob cancels an active fine-tuning job in OpenAI by its ID.
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

// ListJobs retrieves a list of fine-tuning jobs from OpenAI and optionally applies local filtering.
func (s *fineTuningService) ListJobs(ctx context.Context, req *ctrsvc.ListJobsRequest) (*ctrsvc.ListJobsResponse, error) {
	result := &ctrsvc.ListJobsResponse{}

	path := "/fine_tuning/jobs" + s.buildPaginationQuery(req.After, req.Limit)
	resp, err := s.client.RequestRaw(ctx, "GET", path, nil)
	if err != nil {
		return result, fmt.Errorf("retrieve list jobs: %w", err)
	}

	var parsed struct {
		Data []*ctrsvc.Job `json:"data"`
	}
	if err := json.Unmarshal(resp, &parsed); err != nil {
		return result, fmt.Errorf("unmarshal list jobs response: %w", err)
	}

	if s.hasListJobsFilter(req) {
		result.Jobs = s.filterListJobs(parsed.Data, req)
		return result, nil
	}

	result.Jobs = parsed.Data
	return result, nil
}

// ListEvents retrieves a list of fine-tuning job events from OpenAI and optionally applies local filtering.
func (s *fineTuningService) ListEvents(ctx context.Context, req *ctrsvc.ListEventsRequest) (*ctrsvc.ListEventsResponse, error) {
	result := &ctrsvc.ListEventsResponse{}

	path := "/fine_tuning/jobs/" + url.PathEscape(req.JobID) + "/events" + s.buildPaginationQuery(req.After, req.Limit)
	resp, err := s.client.RequestRaw(ctx, "GET", path, nil)
	if err != nil {
		return result, fmt.Errorf("retrieve list events: %w", err)
	}

	var parsed struct {
		Data []*ctrsvc.Event `json:"data"`
	}
	if err := json.Unmarshal(resp, &parsed); err != nil {
		return result, fmt.Errorf("unmarshal list events response: %w", err)
	}

	if s.hasListEventsFilter(req) {
		result.Events = s.filterListEvents(parsed.Data, req)
		return result, nil
	}

	result.Events = parsed.Data
	return result, nil
}

// ListCheckpoints retrieves a list of fine-tuning job checkpoints from OpenAI and optionally applies local filtering.
func (s *fineTuningService) ListCheckpoints(ctx context.Context, req *ctrsvc.ListCheckpointsRequest) (*ctrsvc.ListCheckpointsResponse, error) {
	result := &ctrsvc.ListCheckpointsResponse{}

	path := "/fine_tuning/jobs/" + url.PathEscape(req.JobID) + "/checkpoints" + s.buildPaginationQuery(req.After, req.Limit)
	resp, err := s.client.RequestRaw(ctx, "GET", path, nil)
	if err != nil {
		return result, fmt.Errorf("retrieve list checkpoints: %w", err)
	}

	var parsed struct {
		Data []*ctrsvc.Checkpoint `json:"data"`
	}
	if err := json.Unmarshal(resp, &parsed); err != nil {
		return result, fmt.Errorf("unmarshal list checkpoints response: %w", err)
	}

	if s.hasListCheckpointsFilter(req) {
		result.Checkpoints = s.filterListCheckpoints(parsed.Data, req)
		return result, nil
	}

	result.Checkpoints = parsed.Data
	return result, nil
}

// buildPaginationQuery constructs the API query string from the pagination parameters.
func (*fineTuningService) buildPaginationQuery(after *string, limit *uint8) string {
	q := query.NewUrlQuery()

	q.SetQueryStringParam("after", after)
	q.SetQueryUint8Param("limit", limit)

	return q.Encode()
}

// filterListJobs applies in-memory filtering logic to a list of fine-tuning jobs based on provided conditions.
func (*fineTuningService) filterListJobs(jobs []*ctrsvc.Job, req *ctrsvc.ListJobsRequest) []*ctrsvc.Job {
	var filtered []*ctrsvc.Job
	for i := range jobs {
		if !filter.MatchDateValue(&jobs[i].CreatedAt, req.CreatedAfter, req.CreatedBefore) {
			continue
		}
		if !filter.MatchDateValue(&jobs[i].UpdatedAt, req.UpdatedAfter, req.UpdatedBefore) {
			continue
		}
		if !filter.MatchDateValue(jobs[i].FinishedAt, req.FinishedAfter, req.FinishedBefore) {
			continue
		}
		if !filter.MatchDateValue(jobs[i].EstimatedFinishAt, req.EstimatedFinishAfter, req.EstimatedFinishBefore) {
			continue
		}
		if !filter.MatchStrValue(&jobs[i].ID, req.JobIDs) {
			continue
		}
		if !filter.MatchStrValue(&jobs[i].OrganizationID, req.OrganizationIDs) {
			continue
		}
		if !filter.MatchStrValue(&jobs[i].Status, req.Statuses) {
			continue
		}
		if !filter.MatchStrValue(jobs[i].Suffix, req.Suffixes) {
			continue
		}
		if !filter.MatchStrValue(&jobs[i].Model, req.Models) {
			continue
		}
		if !filter.MatchStrValue(jobs[i].FineTunedModel, req.FineTunedModels) {
			continue
		}
		if !filter.MatchStrValue(&jobs[i].TrainingFile, req.TrainingFiles) {
			continue
		}
		if !filter.MatchStrValue(jobs[i].ValidationFile, req.ValidationFiles) {
			continue
		}
		filtered = append(filtered, jobs[i])
	}
	return filtered
}

// hasListJobsFilter checks whether any of the local filter fields are non-empty or set.
func (*fineTuningService) hasListJobsFilter(req *ctrsvc.ListJobsRequest) bool {
	return len(req.JobIDs) > 0 || len(req.OrganizationIDs) > 0 || len(req.Statuses) > 0 ||
		len(req.Suffixes) > 0 || len(req.Models) > 0 || len(req.FineTunedModels) > 0 ||
		len(req.TrainingFiles) > 0 || len(req.ValidationFiles) > 0 ||
		(req.CreatedAfter != nil && *req.CreatedAfter > 0) ||
		(req.CreatedBefore != nil && *req.CreatedBefore > 0) ||
		(req.UpdatedAfter != nil && *req.UpdatedAfter > 0) ||
		(req.UpdatedBefore != nil && *req.UpdatedBefore > 0) ||
		(req.FinishedAfter != nil && *req.FinishedAfter > 0) ||
		(req.FinishedBefore != nil && *req.FinishedBefore > 0) ||
		(req.EstimatedFinishAfter != nil && *req.EstimatedFinishAfter > 0) ||
		(req.EstimatedFinishBefore != nil && *req.EstimatedFinishBefore > 0)
}

// filterListEvents applies in-memory filtering logic to a list of fine-tuning job events based on provided conditions.
func (*fineTuningService) filterListEvents(events []*ctrsvc.Event, req *ctrsvc.ListEventsRequest) []*ctrsvc.Event {
	var filtered []*ctrsvc.Event
	for i := range events {
		if !filter.MatchDateValue(&events[i].CreatedAt, req.CreatedAfter, req.CreatedBefore) {
			continue
		}
		if !filter.MatchStrValue(&events[i].ID, req.EventIDs) {
			continue
		}
		if !filter.MatchStrValue(&events[i].Level, req.Levels) {
			continue
		}
		filtered = append(filtered, events[i])
	}
	return filtered
}

// hasListEventsFilter checks whether any of the local filter fields are non-empty or set.
func (*fineTuningService) hasListEventsFilter(req *ctrsvc.ListEventsRequest) bool {
	return len(req.EventIDs) > 0 || len(req.Levels) > 0 ||
		(req.CreatedAfter != nil && *req.CreatedAfter > 0) ||
		(req.CreatedBefore != nil && *req.CreatedBefore > 0)
}

// filterListCheckpoints applies in-memory filtering logic to a list of fine-tuning job checkpoints based on provided conditions.
func (*fineTuningService) filterListCheckpoints(checkpoints []*ctrsvc.Checkpoint, req *ctrsvc.ListCheckpointsRequest) []*ctrsvc.Checkpoint {
	var filtered []*ctrsvc.Checkpoint
	for i := range checkpoints {
		if !filter.MatchDateValue(&checkpoints[i].CreatedAt, req.CreatedAfter, req.CreatedBefore) {
			continue
		}
		if !filter.MatchStrValue(&checkpoints[i].ID, req.CheckpointIDs) {
			continue
		}
		filtered = append(filtered, checkpoints[i])
	}
	return filtered
}

// hasListCheckpointsFilter checks whether any of the local filter fields are non-empty or set.
func (*fineTuningService) hasListCheckpointsFilter(req *ctrsvc.ListCheckpointsRequest) bool {
	return len(req.CheckpointIDs) > 0 ||
		(req.CreatedAfter != nil && *req.CreatedAfter > 0) ||
		(req.CreatedBefore != nil && *req.CreatedBefore > 0)
}
