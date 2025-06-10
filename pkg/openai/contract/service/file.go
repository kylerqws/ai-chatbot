package service

import "context"

// File represents a file object returned from the OpenAI API.
type File struct {
	ID        string `json:"id"`
	Object    string `json:"object"`
	Purpose   string `json:"purpose"`
	Filename  string `json:"filename"`
	Bytes     int64  `json:"bytes"`
	CreatedAt int64  `json:"created_at"`
	ExpiresAt *int64 `json:"expires_at,omitempty"`
}

// UploadFileRequest contains parameters for uploading a file to OpenAI.
type UploadFileRequest struct {
	Purpose  string `json:"purpose"`
	FilePath string `json:"file_path"`
}

// UploadFileResponse wraps the uploaded file returned from the API.
type UploadFileResponse struct {
	File *File `json:"file"`
}

// RetrieveFileRequest contains the ID of the file to retrieve.
type RetrieveFileRequest struct {
	FileID string `json:"file_id"`
}

// RetrieveFileResponse wraps the file metadata returned from the API.
type RetrieveFileResponse struct {
	File *File `json:"file"`
}

// RetrieveFileContentRequest contains the ID of the file whose content is to be retrieved.
type RetrieveFileContentRequest struct {
	FileID string `json:"file_id"`
}

// RetrieveFileContentResponse wraps the raw file content as bytes.
type RetrieveFileContentResponse struct {
	Content []byte `json:"content"`
}

// ListFilesRequest contains parameters for filtering listed files.
type ListFilesRequest struct {
	// Local filtering (applied after fetching data)
	FileIDs       []string `json:"file_ids,omitempty"`
	Purposes      []string `json:"purposes,omitempty"`
	Filenames     []string `json:"filenames,omitempty"`
	CreatedAfter  *int64   `json:"created_after,omitempty"`
	CreatedBefore *int64   `json:"created_before,omitempty"`
	ExpiresAfter  *int64   `json:"expires_after,omitempty"`
	ExpiresBefore *int64   `json:"expires_before,omitempty"`

	// API-supported query parameters
	Purpose *string `json:"purpose,omitempty"`
	Order   *string `json:"order,omitempty"`
	After   *string `json:"after,omitempty"`
	Limit   *uint8  `json:"limit,omitempty"`
}

// ListFilesResponse wraps a list of files returned from the API.
type ListFilesResponse struct {
	Files []*File `json:"files"`
}

// DeleteFileRequest contains the ID of the file to delete.
type DeleteFileRequest struct {
	FileID string `json:"file_id"`
}

// DeleteFileResponse contains metadata about the deleted file.
type DeleteFileResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Deleted bool   `json:"deleted"`
}

// FileService defines operations for managing OpenAI files.
type FileService interface {
	// UploadFile uploads a new file to OpenAI.
	UploadFile(ctx context.Context, req *UploadFileRequest) (*UploadFileResponse, error)

	// RetrieveFile retrieves a metadata file by its ID.
	RetrieveFile(ctx context.Context, req *RetrieveFileRequest) (*RetrieveFileResponse, error)

	// RetrieveFileContent retrieves the binary content of a file by its ID.
	RetrieveFileContent(ctx context.Context, req *RetrieveFileContentRequest) (*RetrieveFileContentResponse, error)

	// ListFiles returns a filtered list of uploaded files.
	ListFiles(ctx context.Context, req *ListFilesRequest) (*ListFilesResponse, error)

	// DeleteFile removes a file from OpenAI storage by its ID.
	DeleteFile(ctx context.Context, req *DeleteFileRequest) (*DeleteFileResponse, error)
}
