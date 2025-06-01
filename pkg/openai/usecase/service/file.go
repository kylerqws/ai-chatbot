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

	result.Files = s.applyListFilesFilter(parsed.Data, req)
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

func (*fileService) hasAnyListFilesFilter(req *ctrsvc.ListFilesRequest) bool {
	return req.CreatedAfter != 0 || req.CreatedBefore != 0 ||
		len(req.FileIDs) > 0 || len(req.Statuses) > 0 ||
		len(req.Purposes) > 0 || len(req.Filenames) > 0
}

func (s *fileService) applyListFilesFilter(files []*ctrsvc.File, req *ctrsvc.ListFilesRequest) []*ctrsvc.File {
	var result []*ctrsvc.File
	if !s.hasAnyListFilesFilter(req) {
		return files
	}

	for i := range files {
		if !filter.MatchDateValue(files[i].CreatedAt, req.CreatedAfter, req.CreatedBefore) {
			continue
		}
		if !filter.MatchStrValue(files[i].ID, req.FileIDs) {
			continue
		}
		if !filter.MatchStrValue(files[i].Status, req.Statuses) {
			continue
		}
		if !filter.MatchStrValue(files[i].Purpose, req.Purposes) {
			continue
		}
		if !filter.MatchStrValue(files[i].Filename, req.Filenames) {
			continue
		}

		result = append(result, files[i])
	}

	return result
}
