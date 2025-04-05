package openai

import (
	"context"
	"fmt"

	"github.com/kylerqws/chatbot/pkg/openai/infrastructure/client"
	"github.com/kylerqws/chatbot/pkg/openai/infrastructure/config"
	"github.com/kylerqws/chatbot/pkg/openai/usecase/service"

	ctr "github.com/kylerqws/chatbot/pkg/openai/contract"
	ctrsrv "github.com/kylerqws/chatbot/pkg/openai/contract/service"
)

type api struct {
	fileService ctrsrv.FileService
}

func New(ctx context.Context) (ctr.OpenAI, error) {
	cfg, err := config.New(ctx)
	if err != nil {
		return nil, fmt.Errorf("openai: config init failed: %w", err)
	}

	cl := client.New(cfg)
	sf := service.NewFactory(cl, cfg)

	return &api{
		fileService: sf.FileService(),
	}, nil
}

func (api *api) FileService() ctrsrv.FileService {
	return api.fileService
}
