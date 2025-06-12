package service

import (
	"context"
	ctrsvc "github.com/kylerqws/chatbot/pkg/openai/contract/service"
)

// ModelService defines operations for managing OpenAI models.
type ModelService interface {
	// NewRetrieveModelRequest creates a new retrieve model request.
	NewRetrieveModelRequest() *ctrsvc.RetrieveModelRequest

	// NewRetrieveModelResponse creates a new retrieve model response.
	NewRetrieveModelResponse() *ctrsvc.RetrieveModelResponse

	// RetrieveModel retrieves a model from OpenAI by ID.
	RetrieveModel(ctx context.Context, req *ctrsvc.RetrieveModelRequest) (*ctrsvc.RetrieveModelResponse, error)

	// NewListModelsRequest creates a new list models request.
	NewListModelsRequest() *ctrsvc.ListModelsRequest

	// NewListModelsResponse creates a new list models response.
	NewListModelsResponse() *ctrsvc.ListModelsResponse

	// ListModels retrieves a list of models from OpenAI.
	ListModels(ctx context.Context, req *ctrsvc.ListModelsRequest) (*ctrsvc.ListModelsResponse, error)

	// NewDeleteModelRequest creates a new delete model request.
	NewDeleteModelRequest() *ctrsvc.DeleteModelRequest

	// NewDeleteModelResponse creates a new delete model response.
	NewDeleteModelResponse() *ctrsvc.DeleteModelResponse

	// DeleteModel removes a model from OpenAI by ID.
	DeleteModel(ctx context.Context, req *ctrsvc.DeleteModelRequest) (*ctrsvc.DeleteModelResponse, error)
}
