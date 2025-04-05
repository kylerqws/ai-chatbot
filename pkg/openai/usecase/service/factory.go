package service

import (
	ctrcfg "github.com/kylerqws/chatbot/pkg/openai/contract/config"
	ctrsrv "github.com/kylerqws/chatbot/pkg/openai/contract/service"
	"github.com/kylerqws/chatbot/pkg/openai/infrastructure/client"
	"github.com/kylerqws/chatbot/pkg/openai/usecase/service/handler"
)

type Factory struct {
	fileService ctrsrv.FileService
}

func NewFactory(cl *client.Client, cfg ctrcfg.Config) *Factory {
	return &Factory{
		fileService: handler.NewFileService(cl, cfg),
	}
}

func (f *Factory) FileService() ctrsrv.FileService {
	return f.fileService
}
