package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/kylerqws/chatbot/pkg/openai/domain/purpose"
	"github.com/kylerqws/chatbot/pkg/openai/utils/filter"
	"github.com/kylerqws/chatbot/pkg/openai/utils/query"

	ctrcl "github.com/kylerqws/chatbot/pkg/openai/contract/client"
	ctrcfg "github.com/kylerqws/chatbot/pkg/openai/contract/config"
	ctrsvc "github.com/kylerqws/chatbot/pkg/openai/contract/service"
)

// fileService implements FileService using OpenAI API client.
type fileService struct {
	config ctrcfg.Config
	client ctrcl.Client
}

// NewFileService creates a new FileService instance.
func NewFileService(cl ctrcl.Client, cfg ctrcfg.Config) ctrsvc.FileService {
	return &fileService{config: cfg, client: cl}
}

// UploadFile uploads a file to OpenAI with the given purpose.
func (s *fileService) UploadFile(ctx context.Context, req *ctrsvc.UploadFileRequest) (*ctrsvc.UploadFileResponse, error) {
	result := &ctrsvc.UploadFileResponse{}

	prp, err := purpose.Resolve(req.Purpose)
	if err != nil {
		return result, fmt.Errorf("resolve purpose: %w", err)
	}

	body := map[string]string{"file": req.FilePath, "purpose": prp.Code}
	resp, err := s.client.RequestMultipart(ctx, "/files", body)
	if err != nil {
		return result, fmt.Errorf("upload file: %w", err)
	}

	var file ctrsvc.File
	if err := json.Unmarshal(resp, &file); err != nil {
		return result, fmt.Errorf("unmarshal upload file response: %w", err)
	}

	result.File = &file
	return result, nil
}

// RetrieveFile retrieves a file from OpenAI by its ID.
func (s *fileService) RetrieveFile(ctx context.Context, req *ctrsvc.RetrieveFileRequest) (*ctrsvc.RetrieveFileResponse, error) {
	result := &ctrsvc.RetrieveFileResponse{}

	path := "/files/" + url.PathEscape(req.FileID)
	resp, err := s.client.RequestRaw(ctx, "GET", path, nil)
	if err != nil {
		return result, fmt.Errorf("retrieve file: %w", err)
	}

	var file ctrsvc.File
	if err := json.Unmarshal(resp, &file); err != nil {
		return result, fmt.Errorf("unmarshal retrieve file response: %w", err)
	}

	result.File = &file
	return result, nil
}

// RetrieveFileContent downloads the binary content of a file from OpenAI by ID.
func (s *fileService) RetrieveFileContent(ctx context.Context, req *ctrsvc.RetrieveFileContentRequest) (*ctrsvc.RetrieveFileContentResponse, error) {
	result := &ctrsvc.RetrieveFileContentResponse{}

	path := "/files/" + url.PathEscape(req.FileID) + "/content"
	resp, err := s.client.RequestRaw(ctx, "GET", path, nil)
	if err != nil {
		return result, fmt.Errorf("retrieve file content: %w", err)
	}

	result.Content = resp
	return result, nil
}

// ListFiles retrieves a list of files from OpenAI and optionally applies local filtering.
func (s *fileService) ListFiles(ctx context.Context, req *ctrsvc.ListFilesRequest) (*ctrsvc.ListFilesResponse, error) {
	result := &ctrsvc.ListFilesResponse{}

	path := "/files" + s.buildListFilesQuery(req)
	resp, err := s.client.RequestRaw(ctx, "GET", path, nil)
	if err != nil {
		return result, fmt.Errorf("retrieve list files: %w", err)
	}

	var parsed struct {
		Data []*ctrsvc.File `json:"data"`
	}
	if err := json.Unmarshal(resp, &parsed); err != nil {
		return result, fmt.Errorf("unmarshal list files response: %w", err)
	}

	if s.hasListFilesFilter(req) {
		result.Files = s.filterListFiles(parsed.Data, req)
		return result, nil
	}

	result.Files = parsed.Data
	return result, nil
}

// DeleteFile deletes a file from OpenAI by its ID.
func (s *fileService) DeleteFile(ctx context.Context, req *ctrsvc.DeleteFileRequest) (*ctrsvc.DeleteFileResponse, error) {
	result := &ctrsvc.DeleteFileResponse{}

	path := "/files/" + url.PathEscape(req.FileID)
	resp, err := s.client.RequestRaw(ctx, "DELETE", path, nil)
	if err != nil {
		return result, fmt.Errorf("delete file: %w", err)
	}

	if err := json.Unmarshal(resp, result); err != nil {
		return result, fmt.Errorf("unmarshal delete file response: %w", err)
	}

	if !result.Deleted {
		return result, fmt.Errorf("file not deleted: %s", result.ID)
	}
	return result, nil
}

// buildListFilesQuery constructs the API query string from the filter parameters.
func (*fileService) buildListFilesQuery(req *ctrsvc.ListFilesRequest) string {
	q := query.NewUrlQuery()

	q.SetQueryStringParam("purpose", req.Purpose)
	q.SetQueryStringParam("order", req.Order)
	q.SetQueryStringParam("after", req.After)
	q.SetQueryUint8Param("limit", req.Limit)

	return q.Encode()
}

// filterListFiles applies in-memory filtering logic to a list of files based on provided conditions.
func (*fileService) filterListFiles(files []*ctrsvc.File, req *ctrsvc.ListFilesRequest) []*ctrsvc.File {
	var filtered []*ctrsvc.File
	for i := range files {
		if !filter.MatchDateValue(&files[i].CreatedAt, req.CreatedAfter, req.CreatedBefore) {
			continue
		}
		if !filter.MatchDateValue(files[i].ExpiresAt, req.ExpiresAfter, req.ExpiresBefore) {
			continue
		}
		if !filter.MatchStrValue(&files[i].ID, req.FileIDs) {
			continue
		}
		if !filter.MatchStrValue(&files[i].Purpose, req.Purposes) {
			continue
		}
		if !filter.MatchStrValue(&files[i].Filename, req.Filenames) {
			continue
		}
		filtered = append(filtered, files[i])
	}
	return filtered
}

// hasListFilesFilter checks whether any of the local filter fields are non-empty or set.
func (*fileService) hasListFilesFilter(req *ctrsvc.ListFilesRequest) bool {
	return len(req.FileIDs) > 0 || len(req.Purposes) > 0 || len(req.Filenames) > 0 ||
		(req.CreatedAfter != nil && *req.CreatedAfter > 0) ||
		(req.CreatedBefore != nil && *req.CreatedBefore > 0) ||
		(req.ExpiresAfter != nil && *req.ExpiresAfter > 0) ||
		(req.ExpiresBefore != nil && *req.ExpiresBefore > 0)
}
