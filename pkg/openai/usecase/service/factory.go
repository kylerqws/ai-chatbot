package service

import (
	"github.com/kylerqws/chatbot/pkg/openai/infrastructure/client"
	"github.com/kylerqws/chatbot/pkg/openai/usecase/service/resource"

	ctrlog "github.com/kylerqws/chatbot/pkg/logger/contract/logger"
	ctrcfg "github.com/kylerqws/chatbot/pkg/openai/contract/config"
	ctrsrv "github.com/kylerqws/chatbot/pkg/openai/contract/service"
)

type Factory struct {
	fileService ctrsrv.FileService
}

func NewFactory(cl *client.Client, cfg ctrcfg.Config, log ctrlog.Logger) *Factory {
	return &Factory{
		fileService: resource.NewFileService(cl, cfg, log),
	}
}

func (f *Factory) FileService() ctrsrv.FileService {
	return f.fileService
}
