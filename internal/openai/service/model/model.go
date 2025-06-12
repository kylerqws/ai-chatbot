package model

import (
	"context"

	ctrint "github.com/kylerqws/chatbot/internal/openai/contract/service"
	ctrpkg "github.com/kylerqws/chatbot/pkg/openai/contract"
	ctrsvc "github.com/kylerqws/chatbot/pkg/openai/contract/service"
)

// service provides operations for managing OpenAI models.
type service struct {
	ctx context.Context
	svc ctrsvc.ModelService
}

// NewService creates a new model service for managing OpenAI models.
func NewService(ctx context.Context, sdk ctrpkg.OpenAI) ctrint.ModelService {
	return &service{ctx: ctx, svc: sdk.ModelService()}
}

// NewRetrieveModelRequest creates a new retrieve model request.
func (s *service) NewRetrieveModelRequest() *ctrsvc.RetrieveModelRequest {
	return &ctrsvc.RetrieveModelRequest{}
}

// NewRetrieveModelResponse creates a new retrieve model response.
func (s *service) NewRetrieveModelResponse() *ctrsvc.RetrieveModelResponse {
	return &ctrsvc.RetrieveModelResponse{}
}

// RetrieveModel retrieves a model from OpenAI by ID.
func (s *service) RetrieveModel(ctx context.Context, req *ctrsvc.RetrieveModelRequest) (*ctrsvc.RetrieveModelResponse, error) {
	return s.svc.RetrieveModel(ctx, req)
}

// NewListModelsRequest creates a new list models request.
func (s *service) NewListModelsRequest() *ctrsvc.ListModelsRequest {
	return &ctrsvc.ListModelsRequest{}
}

// NewListModelsResponse creates a new list models response.
func (s *service) NewListModelsResponse() *ctrsvc.ListModelsResponse {
	return &ctrsvc.ListModelsResponse{}
}

// ListModels retrieves a list of models from OpenAI.
func (s *service) ListModels(ctx context.Context, req *ctrsvc.ListModelsRequest) (*ctrsvc.ListModelsResponse, error) {
	return s.svc.ListModels(ctx, req)
}

// NewDeleteModelRequest creates a new delete model request.
func (s *service) NewDeleteModelRequest() *ctrsvc.DeleteModelRequest {
	return &ctrsvc.DeleteModelRequest{}
}

// NewDeleteModelResponse creates a new delete model response.
func (s *service) NewDeleteModelResponse() *ctrsvc.DeleteModelResponse {
	return &ctrsvc.DeleteModelResponse{}
}

// DeleteModel removes a model from OpenAI by ID.
func (s *service) DeleteModel(ctx context.Context, req *ctrsvc.DeleteModelRequest) (*ctrsvc.DeleteModelResponse, error) {
	return s.svc.DeleteModel(ctx, req)
}
