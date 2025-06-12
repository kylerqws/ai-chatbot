package service

import (
	"context"
	"sync"

	"github.com/kylerqws/chatbot/internal/openai/service/chat"
	"github.com/kylerqws/chatbot/internal/openai/service/file"
	"github.com/kylerqws/chatbot/internal/openai/service/fine_tuning"
	"github.com/kylerqws/chatbot/internal/openai/service/model"

	ctrprv "github.com/kylerqws/chatbot/internal/openai/contract/provider"
	ctrsvc "github.com/kylerqws/chatbot/internal/openai/contract/service"
	ctrpkg "github.com/kylerqws/chatbot/pkg/openai/contract"
)

// provider provides access to OpenAI services.
type provider struct {
	ctx context.Context
	sdk ctrpkg.OpenAI

	chat     ctrsvc.ChatService
	chatOnce sync.Once

	file     ctrsvc.FileService
	fileOnce sync.Once

	fineTuning     ctrsvc.FineTuningService
	fineTuningOnce sync.Once

	model     ctrsvc.ModelService
	modelOnce sync.Once
}

// NewProvider creates a new service provider that groups OpenAI API services.
func NewProvider(ctx context.Context, sdk ctrpkg.OpenAI) ctrprv.ServiceProvider {
	return &provider{ctx: ctx, sdk: sdk}
}

// Chat returns the service for chat completions.
func (p *provider) Chat() ctrsvc.ChatService {
	p.chatOnce.Do(func() {
		p.chat = chat.NewService(p.ctx, p.sdk)
	})
	return p.chat
}

// File returns the service for file management.
func (p *provider) File() ctrsvc.FileService {
	p.fileOnce.Do(func() {
		p.file = file.NewService(p.ctx, p.sdk)
	})
	return p.file
}

// FineTuning returns the service for fine-tuning jobs.
func (p *provider) FineTuning() ctrsvc.FineTuningService {
	p.fineTuningOnce.Do(func() {
		p.fineTuning = fine_tuning.NewService(p.ctx, p.sdk)
	})
	return p.fineTuning
}

// Model returns the service for model operations.
func (p *provider) Model() ctrsvc.ModelService {
	p.modelOnce.Do(func() {
		p.model = model.NewService(p.ctx, p.sdk)
	})
	return p.model
}
