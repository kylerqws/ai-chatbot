package resource

import (
	"context"
	"encoding/json"

	"github.com/kylerqws/chatbot/pkg/openai/domain/purpose"
	"github.com/kylerqws/chatbot/pkg/openai/infrastructure/client"

	ctrlog "github.com/kylerqws/chatbot/pkg/logger/contract/logger"
	ctrcfg "github.com/kylerqws/chatbot/pkg/openai/contract/config"
	ctrsrv "github.com/kylerqws/chatbot/pkg/openai/contract/service"
)

type fileService struct {
	config ctrcfg.Config
	logger ctrlog.Logger
	client *client.Client
}

func NewFileService(cl *client.Client, cfg ctrcfg.Config, log ctrlog.Logger) ctrsrv.FileService {
	return &fileService{config: cfg, logger: log, client: cl}
}

func (f *fileService) UploadFile(ctx context.Context, req *ctrsrv.UploadFileRequest) (string, error) {
	prp, err := purpose.Resolve(req.Purpose)
	if err != nil {
		return "", err
	}

	resp, err := f.client.RequestMultipart(ctx, "/v1/files", map[string]string{
		"file":    req.FilePath,
		"purpose": prp.Value,
	}, "file")
	if err != nil {
		return "", err
	}

	var result struct {
		ID string `json:"id"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return "", err
	}

	return result.ID, nil
}

func (f *fileService) GetFileInfo(ctx context.Context, req *ctrsrv.GetFileInfoRequest) (*ctrsrv.File, error) {
	resp, err := f.client.Request(ctx, "GET", "/v1/files/"+req.FileID)
	if err != nil {
		return nil, err
	}

	var result ctrsrv.File
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (f *fileService) ListFiles(ctx context.Context, _ *ctrsrv.ListFilesRequest) ([]*ctrsrv.File, error) {
	resp, err := f.client.Request(ctx, "GET", "/v1/files")
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []*ctrsrv.File `json:"data"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (f *fileService) DeleteFile(ctx context.Context, req *ctrsrv.DeleteFileRequest) error {
	_, err := f.client.Request(ctx, "DELETE", "/v1/files/"+req.FileID)
	return err
}
