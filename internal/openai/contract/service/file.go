package service

import (
	"context"
	ctrsvc "github.com/kylerqws/chatbot/pkg/openai/contract/service"
)

// FileService defines operations for managing OpenAI files.
type FileService interface {
	// NewUploadFileRequest creates a new upload file request.
	NewUploadFileRequest() *ctrsvc.UploadFileRequest

	// NewUploadFileResponse creates a new upload file response.
	NewUploadFileResponse() *ctrsvc.UploadFileResponse

	// UploadFile uploads a file to OpenAI.
	UploadFile(ctx context.Context, req *ctrsvc.UploadFileRequest) (*ctrsvc.UploadFileResponse, error)

	// NewRetrieveFileRequest creates a new retrieve file request.
	NewRetrieveFileRequest() *ctrsvc.RetrieveFileRequest

	// NewRetrieveFileResponse creates a new retrieve file response.
	NewRetrieveFileResponse() *ctrsvc.RetrieveFileResponse

	// RetrieveFile retrieves file metadata from OpenAI by ID.
	RetrieveFile(ctx context.Context, req *ctrsvc.RetrieveFileRequest) (*ctrsvc.RetrieveFileResponse, error)

	// NewRetrieveFileContentRequest creates a new retrieve file content request.
	NewRetrieveFileContentRequest() *ctrsvc.RetrieveFileContentRequest

	// NewRetrieveFileContentResponse creates a new retrieve file content response.
	NewRetrieveFileContentResponse() *ctrsvc.RetrieveFileContentResponse

	// RetrieveFileContent retrieves the binary content of a file from OpenAI by ID.
	RetrieveFileContent(ctx context.Context, req *ctrsvc.RetrieveFileContentRequest) (*ctrsvc.RetrieveFileContentResponse, error)

	// NewListFilesRequest creates a new list files request.
	NewListFilesRequest() *ctrsvc.ListFilesRequest

	// NewListFilesResponse creates a new list files response.
	NewListFilesResponse() *ctrsvc.ListFilesResponse

	// ListFiles retrieves a list of files from OpenAI.
	ListFiles(ctx context.Context, req *ctrsvc.ListFilesRequest) (*ctrsvc.ListFilesResponse, error)

	// NewDeleteFileRequest creates a new delete file request.
	NewDeleteFileRequest() *ctrsvc.DeleteFileRequest

	// NewDeleteFileResponse creates a new delete file response.
	NewDeleteFileResponse() *ctrsvc.DeleteFileResponse

	// DeleteFile removes a file from OpenAI by ID.
	DeleteFile(ctx context.Context, req *ctrsvc.DeleteFileRequest) (*ctrsvc.DeleteFileResponse, error)
}
