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

	chatServiceOnce sync.Once
	chatService     ctrsvc.ChatService

	fileServiceOnce sync.Once
	fileService     ctrsvc.FileService

	fineServiceOnce   sync.Once
	fineTuningService ctrsvc.FineTuningService

	modelServiceOnce sync.Once
	modelService     ctrsvc.ModelService
}

// NewProvider creates a new service provider that groups OpenAI API services.
func NewProvider(ctx context.Context, sdk ctrpkg.OpenAI) ctrprv.ServiceProvider {
	return &provider{ctx: ctx, sdk: sdk}
}

// Chat returns the service for chat completions.
func (p *provider) Chat() ctrsvc.ChatService {
	p.chatServiceOnce.Do(func() {
		p.chatService = chat.NewService(p.ctx, p.sdk)
	})
	return p.chatService
}

// File returns the service for file management.
func (p *provider) File() ctrsvc.FileService {
	p.fileServiceOnce.Do(func() {
		p.fileService = file.NewService(p.ctx, p.sdk)
	})
	return p.fileService
}

// FineTuning returns the service for fine-tuning jobs.
func (p *provider) FineTuning() ctrsvc.FineTuningService {
	p.fineServiceOnce.Do(func() {
		p.fineTuningService = fine_tuning.NewService(p.ctx, p.sdk)
	})
	return p.fineTuningService
}

// Model returns the service for model operations.
func (p *provider) Model() ctrsvc.ModelService {
	p.modelServiceOnce.Do(func() {
		p.modelService = model.NewService(p.ctx, p.sdk)
	})
	return p.modelService
}
