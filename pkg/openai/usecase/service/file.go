package service

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kylerqws/chatbot/pkg/openai/domain/purpose"
	"github.com/kylerqws/chatbot/pkg/openai/infrastructure/client"
	"github.com/kylerqws/chatbot/pkg/openai/utils/filter"

	ctrcfg "github.com/kylerqws/chatbot/pkg/openai/contract/config"
	ctrsvc "github.com/kylerqws/chatbot/pkg/openai/contract/service"
)

type fileService struct {
	config ctrcfg.Config
	client *client.Client
}

func NewFileService(cl *client.Client, cfg ctrcfg.Config) ctrsvc.FileService {
	return &fileService{config: cfg, client: cl}
}

func (s *fileService) UploadFile(
	ctx context.Context,
	req *ctrsvc.UploadFileRequest,
) (*ctrsvc.UploadFileResponse, error) {
	result := &ctrsvc.UploadFileResponse{}

	prp, err := purpose.Resolve(req.Purpose)
	if err != nil {
		return result, fmt.Errorf("failed to resolve purpose: %w", err)
	}

	body := map[string]string{"file": req.FilePath, "purpose": prp.Code}
	resp, err := s.client.RequestMultipart(ctx, "/files", body)
	if err != nil {
		return result, fmt.Errorf("failed to send multipart request: %w", err)
	}

	var file ctrsvc.File
	err = json.Unmarshal(resp, &file)
	if err != nil {
		return result, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	result.File = &file
	return result, nil
}

func (s *fileService) GetFileInfo(
	ctx context.Context,
	req *ctrsvc.GetFileInfoRequest,
) (*ctrsvc.GetFileInfoResponse, error) {
	result := &ctrsvc.GetFileInfoResponse{}

	resp, err := s.client.Request(ctx, "GET", "/files/"+req.FileID)
	if err != nil {
		return result, fmt.Errorf("failed to send request: %w", err)
	}

	var file ctrsvc.File
	err = json.Unmarshal(resp, &file)
	if err != nil {
		return result, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	result.File = &file
	return result, nil
}

func (s *fileService) ListFiles(
	ctx context.Context,
	req *ctrsvc.ListFilesRequest,
) (*ctrsvc.ListFilesResponse, error) {
	result := &ctrsvc.ListFilesResponse{}

	resp, err := s.client.Request(ctx, "GET", "/files")
	if err != nil {
		return result, fmt.Errorf("failed to send request: %w", err)
	}

	var parsed struct {
		Data []*ctrsvc.File `json:"data"`
	}
	err = json.Unmarshal(resp, &parsed)
	if err != nil {
		return result, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	result.Files = s.filterFiles(parsed.Data, req)
	return result, nil
}

func (s *fileService) DeleteFile(
	ctx context.Context,
	req *ctrsvc.DeleteFileRequest,
) (*ctrsvc.DeleteFileResponse, error) {
	result := &ctrsvc.DeleteFileResponse{}

	resp, err := s.client.Request(ctx, "DELETE", "/files/"+req.FileID)
	if err != nil {
		return result, fmt.Errorf("failed to send request: %w", err)
	}

	err = json.Unmarshal(resp, result)
	if err != nil {
		return result, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if !result.Deleted {
		return result, fmt.Errorf("failed to delete file '%v'", result.ID)
	}
	return result, nil
}

func (s *fileService) filterFiles(files []*ctrsvc.File, req *ctrsvc.ListFilesRequest) []*ctrsvc.File {
	var result []*ctrsvc.File

	for i := range files {
		if filter.CheckDateValue(files[i].CreatedAt, req.CreatedAfter, req.CreatedBefore) {
			continue
		}
		if filter.CheckStrValue(files[i].ID, req.FileIDs) {
			continue
		}
		if filter.CheckStrValue(files[i].Status, req.Statuses) {
			continue
		}
		if filter.CheckStrValue(files[i].Purpose, req.Purposes) {
			continue
		}
		if filter.CheckStrValue(files[i].Filename, req.Filenames) {
			continue
		}

		result = append(result, files[i])
	}

	return result
}
