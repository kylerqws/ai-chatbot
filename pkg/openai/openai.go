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

// openai aggregates services for working with the OpenAI API.
type openai struct {
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
	return &openai{ctx: ctx, cfg: cfg}, nil
}

// FileService returns the initialized FileService client.
func (o *openai) FileService() ctrsvc.FileService {
	o.fileOnce.Do(func() {
		o.file = service.NewFileService(o.client(), o.cfg)
	})
	return o.file
}

// FineTuningService returns the initialized FineTuningService client.
func (o *openai) FineTuningService() ctrsvc.FineTuningService {
	o.fineTuningOnce.Do(func() {
		o.fineTuning = service.NewFineTuningService(o.client(), o.cfg)
	})
	return o.fineTuning
}

// ModelService returns the initialized ModelService client.
func (o *openai) ModelService() ctrsvc.ModelService {
	o.modelOnce.Do(func() {
		o.model = service.NewModelService(o.client(), o.cfg)
	})
	return o.model
}

// ChatService returns the initialized ChatService client.
func (o *openai) ChatService() ctrsvc.ChatService {
	o.chatOnce.Do(func() {
		o.chat = service.NewChatService(o.client(), o.cfg)
	})
	return o.chat
}

// client returns the initialized OpenAI HTTP client.
func (o *openai) client() ctrcl.Client {
	o.clOnce.Do(func() {
		o.cl = client.New(o.cfg)
	})
	return o.cl
}
