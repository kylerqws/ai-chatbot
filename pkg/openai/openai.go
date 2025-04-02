package openai

import (
	"context"

	"github.com/kylerqws/chatbot/pkg/openai/infrastructure/client"
	"github.com/kylerqws/chatbot/pkg/openai/infrastructure/config"
	"github.com/kylerqws/chatbot/pkg/openai/usecase/service"

	ctrlog "github.com/kylerqws/chatbot/pkg/logger/contract/logger"
	ctrsrv "github.com/kylerqws/chatbot/pkg/openai/contract/service"
)

type API struct {
	fileService ctrsrv.FileService
}

func New(ctx context.Context, log ctrlog.Logger) (*API, error) {
	cfg, err := config.New(ctx)
	if err != nil {
		return nil, err
	}

	cl := client.New(cfg, log)
	sf := service.NewFactory(cl, cfg, log)

	return &API{
		fileService: sf.FileService(),
	}, nil
}

func (api *API) FileService() ctrsrv.FileService {
	return api.fileService
}
