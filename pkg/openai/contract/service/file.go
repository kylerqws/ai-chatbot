package service

import "context"

type File struct {
	ID        string `json:"id"`
	Object    string `json:"object"`
	Bytes     int64  `json:"bytes"`
	CreatedAt int64  `json:"created_at"`
	Filename  string `json:"filename"`
	Purpose   string `json:"purpose"`
	Status    string `json:"status,omitempty"`
}

type UploadFileRequest struct {
	FilePath string `json:"file_path"`
	Purpose  string `json:"purpose"`
}

type UploadFileResponse struct {
	File *File `json:"file"`
}

type GetFileInfoRequest struct {
	FileID string `json:"file_id"`
}

type GetFileInfoResponse struct {
	File *File `json:"file"`
}

type ListFilesRequest struct {
	FileIDs       []string `json:"file_ids,omitempty"`
	Purposes      []string `json:"purposes,omitempty"`
	Filenames     []string `json:"filenames,omitempty"`
	Statuses      []string `json:"statuses,omitempty"`
	CreatedAfter  int64    `json:"created_after,omitempty"`
	CreatedBefore int64    `json:"created_before,omitempty"`
}

type ListFilesResponse struct {
	Files []*File `json:"files"`
}

type DeleteFileRequest struct {
	FileID string `json:"file_id"`
}

type DeleteFileResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Deleted bool   `json:"deleted"`
}

type FileService interface {
	UploadFile(ctx context.Context, req *UploadFileRequest) (*UploadFileResponse, error)
	GetFileInfo(ctx context.Context, req *GetFileInfoRequest) (*GetFileInfoResponse, error)
	ListFiles(ctx context.Context, req *ListFilesRequest) (*ListFilesResponse, error)
	DeleteFile(ctx context.Context, req *DeleteFileRequest) (*DeleteFileResponse, error)
}
