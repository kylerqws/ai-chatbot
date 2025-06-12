package file

import (
	"context"

	ctrint "github.com/kylerqws/chatbot/internal/openai/contract/service"
	ctrpkg "github.com/kylerqws/chatbot/pkg/openai/contract"
	ctrsvc "github.com/kylerqws/chatbot/pkg/openai/contract/service"
)

// service provides operations for managing OpenAI files.
type service struct {
	ctx context.Context
	svc ctrsvc.FileService
}

// NewService creates a new file service for managing OpenAI files.
func NewService(ctx context.Context, sdk ctrpkg.OpenAI) ctrint.FileService {
	return &service{ctx: ctx, svc: sdk.FileService()}
}

// NewUploadFileRequest creates a new upload file request.
func (s *service) NewUploadFileRequest() *ctrsvc.UploadFileRequest {
	return &ctrsvc.UploadFileRequest{}
}

// NewUploadFileResponse creates a new upload file response.
func (s *service) NewUploadFileResponse() *ctrsvc.UploadFileResponse {
	return &ctrsvc.UploadFileResponse{}
}

// UploadFile uploads a file to OpenAI.
func (s *service) UploadFile(ctx context.Context, req *ctrsvc.UploadFileRequest) (*ctrsvc.UploadFileResponse, error) {
	return s.svc.UploadFile(ctx, req)
}

// NewRetrieveFileRequest creates a new retrieve file request.
func (s *service) NewRetrieveFileRequest() *ctrsvc.RetrieveFileRequest {
	return &ctrsvc.RetrieveFileRequest{}
}

// NewRetrieveFileResponse creates a new retrieve file response.
func (s *service) NewRetrieveFileResponse() *ctrsvc.RetrieveFileResponse {
	return &ctrsvc.RetrieveFileResponse{}
}

// RetrieveFile retrieves file metadata from OpenAI by ID.
func (s *service) RetrieveFile(ctx context.Context, req *ctrsvc.RetrieveFileRequest) (*ctrsvc.RetrieveFileResponse, error) {
	return s.svc.RetrieveFile(ctx, req)
}

// NewRetrieveFileContentRequest creates a new retrieve file content request.
func (s *service) NewRetrieveFileContentRequest() *ctrsvc.RetrieveFileContentRequest {
	return &ctrsvc.RetrieveFileContentRequest{}
}

// NewRetrieveFileContentResponse creates a new retrieve file content response.
func (s *service) NewRetrieveFileContentResponse() *ctrsvc.RetrieveFileContentResponse {
	return &ctrsvc.RetrieveFileContentResponse{}
}

// RetrieveFileContent retrieves the binary content of a file from OpenAI by ID.
func (s *service) RetrieveFileContent(ctx context.Context, req *ctrsvc.RetrieveFileContentRequest) (*ctrsvc.RetrieveFileContentResponse, error) {
	return s.svc.RetrieveFileContent(ctx, req)
}

// NewListFilesRequest creates a new list files request.
func (s *service) NewListFilesRequest() *ctrsvc.ListFilesRequest {
	return &ctrsvc.ListFilesRequest{}
}

// NewListFilesResponse creates a new list files response.
func (s *service) NewListFilesResponse() *ctrsvc.ListFilesResponse {
	return &ctrsvc.ListFilesResponse{}
}

// ListFiles retrieves a list of files from OpenAI.
func (s *service) ListFiles(ctx context.Context, req *ctrsvc.ListFilesRequest) (*ctrsvc.ListFilesResponse, error) {
	return s.svc.ListFiles(ctx, req)
}

// NewDeleteFileRequest creates a new delete file request.
func (s *service) NewDeleteFileRequest() *ctrsvc.DeleteFileRequest {
	return &ctrsvc.DeleteFileRequest{}
}

// NewDeleteFileResponse creates a new delete file response.
func (s *service) NewDeleteFileResponse() *ctrsvc.DeleteFileResponse {
	return &ctrsvc.DeleteFileResponse{}
}

// DeleteFile removes a file from OpenAI by ID.
func (s *service) DeleteFile(ctx context.Context, req *ctrsvc.DeleteFileRequest) (*ctrsvc.DeleteFileResponse, error) {
	return s.svc.DeleteFile(ctx, req)
}
