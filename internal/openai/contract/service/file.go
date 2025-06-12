package service

import (
	"context"
	ctrsvc "github.com/kylerqws/chatbot/pkg/openai/contract/service"
)

// FileService defines operations for managing OpenAI files.
type FileService interface {
	// NewUploadFileRequest returns a new upload file request instance.
	NewUploadFileRequest() *ctrsvc.UploadFileRequest

	// NewUploadFileResponse returns a new upload file response instance.
	NewUploadFileResponse() *ctrsvc.UploadFileResponse

	// UploadFile uploads a file to OpenAI.
	UploadFile(ctx context.Context, req *ctrsvc.UploadFileRequest) (*ctrsvc.UploadFileResponse, error)

	// NewRetrieveFileRequest returns a new retrieve file request instance.
	NewRetrieveFileRequest() *ctrsvc.RetrieveFileRequest

	// NewRetrieveFileResponse returns a new retrieve file response instance.
	NewRetrieveFileResponse() *ctrsvc.RetrieveFileResponse

	// RetrieveFile retrieves file metadata from OpenAI by ID.
	RetrieveFile(ctx context.Context, req *ctrsvc.RetrieveFileRequest) (*ctrsvc.RetrieveFileResponse, error)

	// NewRetrieveFileContentRequest returns a new retrieve file content request instance.
	NewRetrieveFileContentRequest() *ctrsvc.RetrieveFileContentRequest

	// NewRetrieveFileContentResponse returns a new retrieve file content response instance.
	NewRetrieveFileContentResponse() *ctrsvc.RetrieveFileContentResponse

	// RetrieveFileContent retrieves the binary content of a file from OpenAI by ID.
	RetrieveFileContent(ctx context.Context, req *ctrsvc.RetrieveFileContentRequest) (*ctrsvc.RetrieveFileContentResponse, error)

	// NewListFilesRequest returns a new list files request instance.
	NewListFilesRequest() *ctrsvc.ListFilesRequest

	// NewListFilesResponse returns a new list files response instance.
	NewListFilesResponse() *ctrsvc.ListFilesResponse

	// ListFiles retrieves a list of files from OpenAI.
	ListFiles(ctx context.Context, req *ctrsvc.ListFilesRequest) (*ctrsvc.ListFilesResponse, error)

	// NewDeleteFileRequest returns a new delete file request instance.
	NewDeleteFileRequest() *ctrsvc.DeleteFileRequest

	// NewDeleteFileResponse returns a new delete file response instance.
	NewDeleteFileResponse() *ctrsvc.DeleteFileResponse

	// DeleteFile removes a file from OpenAI by ID.
	DeleteFile(ctx context.Context, req *ctrsvc.DeleteFileRequest) (*ctrsvc.DeleteFileResponse, error)
}
