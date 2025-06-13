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

// entrypoint aggregates access to OpenAI API services.
type entrypoint struct {
	ctx context.Context
	cfg ctrcfg.Config

	cl     ctrcl.Client
	clOnce sync.Once

	file     ctrsvc.FileService
	fileOnce sync.Once

	fineTuning     ctrsvc.FineTuningService
	fineTuningOnce sync.Once

	model     ctrsvc.ModelService
	modelOnce sync.Once

	chat     ctrsvc.ChatService
	chatOnce sync.Once
}

// New creates and returns a new OpenAI API access object.
func New(ctx context.Context) (ctr.OpenAI, error) {
	cfg, err := config.New(ctx)
	if err != nil {
		return nil, fmt.Errorf("load openai config: %w", err)
	}
	return &entrypoint{ctx: ctx, cfg: cfg}, nil
}

// FileService returns the initialized FileService client.
func (e *entrypoint) FileService() ctrsvc.FileService {
	e.fileOnce.Do(func() {
		e.file = service.NewFileService(e.client(), e.cfg)
	})
	return e.file
}

// FineTuningService returns the initialized FineTuningService client.
func (e *entrypoint) FineTuningService() ctrsvc.FineTuningService {
	e.fineTuningOnce.Do(func() {
		e.fineTuning = service.NewFineTuningService(e.client(), e.cfg)
	})
	return e.fineTuning
}

// ModelService returns the initialized ModelService client.
func (e *entrypoint) ModelService() ctrsvc.ModelService {
	e.modelOnce.Do(func() {
		e.model = service.NewModelService(e.client(), e.cfg)
	})
	return e.model
}

// ChatService returns the initialized ChatService client.
func (e *entrypoint) ChatService() ctrsvc.ChatService {
	e.chatOnce.Do(func() {
		e.chat = service.NewChatService(e.client(), e.cfg)
	})
	return e.chat
}

// client returns the initialized OpenAI HTTP client.
func (e *entrypoint) client() ctrcl.Client {
	e.clOnce.Do(func() {
		e.cl = client.New(e.cfg)
	})
	return e.cl
}
