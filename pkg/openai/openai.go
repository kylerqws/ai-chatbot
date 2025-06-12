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

// manager aggregates services for working with the OpenAI API.
type manager struct {
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
	return &manager{ctx: ctx, cfg: cfg}, nil
}

// FileService returns the initialized FileService client.
func (m *manager) FileService() ctrsvc.FileService {
	m.fileOnce.Do(func() {
		m.file = service.NewFileService(m.client(), m.cfg)
	})
	return m.file
}

// FineTuningService returns the initialized FineTuningService client.
func (m *manager) FineTuningService() ctrsvc.FineTuningService {
	m.fineTuningOnce.Do(func() {
		m.fineTuning = service.NewFineTuningService(m.client(), m.cfg)
	})
	return m.fineTuning
}

// ModelService returns the initialized ModelService client.
func (m *manager) ModelService() ctrsvc.ModelService {
	m.modelOnce.Do(func() {
		m.model = service.NewModelService(m.client(), m.cfg)
	})
	return m.model
}

// ChatService returns the initialized ChatService client.
func (m *manager) ChatService() ctrsvc.ChatService {
	m.chatOnce.Do(func() {
		m.chat = service.NewChatService(m.client(), m.cfg)
	})
	return m.chat
}

// client returns the initialized OpenAI HTTP client.
func (m *manager) client() ctrcl.Client {
	m.clOnce.Do(func() {
		m.cl = client.New(m.cfg)
	})
	return m.cl
}
