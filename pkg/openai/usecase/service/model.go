package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/kylerqws/chatbot/pkg/openai/utils/filter"

	ctrcl "github.com/kylerqws/chatbot/pkg/openai/contract/client"
	ctrcfg "github.com/kylerqws/chatbot/pkg/openai/contract/config"
	ctrsvc "github.com/kylerqws/chatbot/pkg/openai/contract/service"
)

// modelService implements ModelService using OpenAI API client.
type modelService struct {
	config ctrcfg.Config
	client ctrcl.Client
}

// NewModelService creates a new ModelService instance.
func NewModelService(cl ctrcl.Client, cfg ctrcfg.Config) ctrsvc.ModelService {
	return &modelService{config: cfg, client: cl}
}

// RetrieveModel retrieves a model from OpenAI by its ID.
func (s *modelService) RetrieveModel(ctx context.Context, req *ctrsvc.RetrieveModelRequest) (*ctrsvc.RetrieveModelResponse, error) {
	result := &ctrsvc.RetrieveModelResponse{}

	path := "/models/" + url.PathEscape(req.ModelID)
	resp, err := s.client.RequestRaw(ctx, "GET", path, nil)
	if err != nil {
		return result, fmt.Errorf("retrieve model: %w", err)
	}

	var model ctrsvc.Model
	if err := json.Unmarshal(resp, &model); err != nil {
		return result, fmt.Errorf("unmarshal retrieve model response: %w", err)
	}

	result.Model = &model
	return result, nil
}

// ListModels retrieves a list of models from OpenAI and optionally applies local filtering.
func (s *modelService) ListModels(ctx context.Context, req *ctrsvc.ListModelsRequest) (*ctrsvc.ListModelsResponse, error) {
	result := &ctrsvc.ListModelsResponse{}

	resp, err := s.client.RequestRaw(ctx, "GET", "/models", nil)
	if err != nil {
		return result, fmt.Errorf("retrieve list models: %w", err)
	}

	var parsed struct {
		Data []*ctrsvc.Model `json:"data"`
	}
	if err := json.Unmarshal(resp, &parsed); err != nil {
		return result, fmt.Errorf("unmarshal list models response: %w", err)
	}

	if s.hasListModelsFilter(req) {
		result.Models = s.filterListModels(parsed.Data, req)
		return result, nil
	}

	result.Models = parsed.Data
	return result, nil
}

// DeleteModel deletes a model from OpenAI by its ID.
func (s *modelService) DeleteModel(ctx context.Context, req *ctrsvc.DeleteModelRequest) (*ctrsvc.DeleteModelResponse, error) {
	result := &ctrsvc.DeleteModelResponse{}

	path := "/models/" + url.PathEscape(req.ModelID)
	resp, err := s.client.RequestRaw(ctx, "DELETE", path, nil)
	if err != nil {
		return result, fmt.Errorf("delete model: %w", err)
	}

	if err := json.Unmarshal(resp, result); err != nil {
		return result, fmt.Errorf("unmarshal delete model response: %w", err)
	}

	if !result.Deleted {
		return result, fmt.Errorf("model not deleted: %s", result.ID)
	}
	return result, nil
}

// filterListModels applies in-memory filtering logic to a list of models based on provided conditions.
func (*modelService) filterListModels(models []*ctrsvc.Model, req *ctrsvc.ListModelsRequest) []*ctrsvc.Model {
	var filtered []*ctrsvc.Model
	for i := range models {
		if !filter.MatchDateValue(&models[i].CreatedAt, req.CreatedAfter, req.CreatedBefore) {
			continue
		}
		if !filter.MatchStrValue(&models[i].ID, req.ModelIDs) {
			continue
		}
		if !filter.MatchStrValue(&models[i].OwnedBy, req.Owners) {
			continue
		}
		filtered = append(filtered, models[i])
	}
	return filtered
}

// hasListModelsFilter checks whether any of the local filter fields are non-empty or set.
func (*modelService) hasListModelsFilter(req *ctrsvc.ListModelsRequest) bool {
	return len(req.ModelIDs) > 0 || len(req.Owners) > 0 ||
		(req.CreatedAfter != nil && *req.CreatedAfter > 0) ||
		(req.CreatedBefore != nil && *req.CreatedBefore > 0)
}
