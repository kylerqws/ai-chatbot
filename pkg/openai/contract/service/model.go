package service

import "context"

// Model represents a model object returned from the OpenAI API.
type Model struct {
	ID        string `json:"id"`
	Object    string `json:"object"`
	OwnedBy   string `json:"owned_by"`
	CreatedAt int64  `json:"created"`
}

// RetrieveModelRequest contains the ID of the model to retrieve.
type RetrieveModelRequest struct {
	ModelID string `json:"model_id"`
}

// RetrieveModelResponse wraps the model metadata returned from the API.
type RetrieveModelResponse struct {
	Model *Model `json:"model"`
}

// ListModelsRequest contains parameters for filtering listed models.
type ListModelsRequest struct {
	// Local filtering (applied after fetching data)
	ModelIDs      []string `json:"model_ids,omitempty"`
	Owners        []string `json:"owners,omitempty"`
	CreatedAfter  *int64   `json:"created_after,omitempty"`
	CreatedBefore *int64   `json:"created_before,omitempty"`
}

// ListModelsResponse wraps a list of models returned from the API.
type ListModelsResponse struct {
	Models []*Model `json:"models"`
}

// DeleteModelRequest contains the ID of the model to delete.
type DeleteModelRequest struct {
	ModelID string `json:"model_id"`
}

// DeleteModelResponse contains metadata about the deleted model.
type DeleteModelResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Deleted bool   `json:"deleted"`
}

// ModelService defines operations for managing OpenAI models.
type ModelService interface {
	// RetrieveModel retrieves a metadata model by its ID.
	RetrieveModel(ctx context.Context, req *RetrieveModelRequest) (*RetrieveModelResponse, error)

	// ListModels returns a filtered list of available models.
	ListModels(ctx context.Context, req *ListModelsRequest) (*ListModelsResponse, error)

	// DeleteModel removes a model from OpenAI by its ID.
	DeleteModel(ctx context.Context, req *DeleteModelRequest) (*DeleteModelResponse, error)
}
