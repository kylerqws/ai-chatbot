package openai

import (
	"context"
	"fmt"
	"sync"

	"github.com/kylerqws/chatbot/pkg/openai/infrastructure/client"
	"github.com/kylerqws/chatbot/pkg/openai/infrastructure/config"
	"github.com/kylerqws/chatbot/pkg/openai/usecase/service"

	ctr "github.com/kylerqws/chatbot/pkg/openai/contract"
	ctrcl "github.com/kylerqws/chatbot/pkg/openai/contract/client"
	ctrcfg "github.com/kylerqws/chatbot/pkg/openai/contract/config"
	ctrsvc "github.com/kylerqws/chatbot/pkg/openai/contract/service"
)

// api aggregates services for working with the OpenAI API.
type api struct {
	ctx context.Context
	cfg ctrcfg.Config

	cl     ctrcl.Client
	clOnce sync.Once

	chat     ctrsvc.ChatService
	chatOnce sync.Once

	file     ctrsvc.FileService
	fileOnce sync.Once

	fineTuning     ctrsvc.FineTuningService
	fineTuningOnce sync.Once

	model     ctrsvc.ModelService
	modelOnce sync.Once
}

// New creates and returns a new OpenAI API access object.
func New(ctx context.Context) (ctr.OpenAI, error) {
	cfg, err := config.New(ctx)
	if err != nil {
		return nil, fmt.Errorf("load openai config: %w", err)
	}
	return &api{ctx: ctx, cfg: cfg}, nil
}

// FileService returns the initialized FileService client.
func (a *api) FileService() ctrsvc.FileService {
	a.fileOnce.Do(func() {
		a.file = service.NewFileService(a.client(), a.cfg)
	})
	return a.file
}

// FineTuningService returns the initialized FineTuningService client.
func (a *api) FineTuningService() ctrsvc.FineTuningService {
	a.fineTuningOnce.Do(func() {
		a.fineTuning = service.NewFineTuningService(a.client(), a.cfg)
	})
	return a.fineTuning
}

// ModelService returns the initialized ModelService client.
func (a *api) ModelService() ctrsvc.ModelService {
	a.modelOnce.Do(func() {
		a.model = service.NewModelService(a.client(), a.cfg)
	})
	return a.model
}

// ChatService returns the initialized ChatService client.
func (a *api) ChatService() ctrsvc.ChatService {
	a.chatOnce.Do(func() {
		a.chat = service.NewChatService(a.client(), a.cfg)
	})
	return a.chat
}

// client returns the initialized OpenAI HTTP client.
func (a *api) client() ctrcl.Client {
	a.clOnce.Do(func() {
		a.cl = client.New(a.cfg)
	})
	return a.cl
}
