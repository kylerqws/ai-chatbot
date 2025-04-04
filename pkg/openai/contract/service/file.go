package service

import "context"

type File struct {
	ID        string `json:"id"`
	Object    string `json:"object"`
	Bytes     int    `json:"bytes"`
	CreatedAt int64  `json:"created_at"`
	Filename  string `json:"filename"`
	Purpose   string `json:"purpose"`
}

type UploadFileRequest struct {
	FilePath string
	Purpose  string
}

type GetFileInfoRequest struct {
	FileID string
}

type DeleteFileRequest struct {
	FileID string
}

type ListFilesRequest struct {
	// TODO: will need to add filters
}

type FileService interface {
	UploadFile(ctx context.Context, req *UploadFileRequest) (string, error)
	GetFileInfo(ctx context.Context, req *GetFileInfoRequest) (*File, error)
	ListFiles(ctx context.Context, req *ListFilesRequest) ([]*File, error)
	DeleteFile(ctx context.Context, req *DeleteFileRequest) error
}
