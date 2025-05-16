package openai

import (
	"context"
	"fmt"

	"github.com/kylerqws/chatbot/pkg/openai/infrastructure/client"
	"github.com/kylerqws/chatbot/pkg/openai/infrastructure/config"
	"github.com/kylerqws/chatbot/pkg/openai/usecase/service"

	ctr "github.com/kylerqws/chatbot/pkg/openai/contract"
	ctrsvc "github.com/kylerqws/chatbot/pkg/openai/contract/service"
)

type api struct {
	fileService ctrsvc.FileService
	jobService  ctrsvc.JobService
	chatService ctrsvc.ChatService
}

func New(ctx context.Context) (ctr.OpenAI, error) {
	cfg, err := config.New(ctx)
	if err != nil {
		return nil, fmt.Errorf("[openai.New] failed to load config: %w", err)
	}

	cl := client.New(cfg)
	return &api{
		fileService: service.NewFileService(cl, cfg),
		jobService:  service.NewJobService(cl, cfg),
		chatService: service.NewChatService(cl, cfg),
	}, nil
}

func (api *api) FileService() ctrsvc.FileService {
	return api.fileService
}

func (api *api) JobService() ctrsvc.JobService {
	return api.jobService
}

func (api *api) ChatService() ctrsvc.ChatService {
	return api.chatService
}
