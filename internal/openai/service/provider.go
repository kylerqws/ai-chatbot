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
func (m *provider) Chat() ctrsvc.ChatService {
	m.chatServiceOnce.Do(func() {
		m.chatService = chat.NewService(m.ctx, m.sdk)
	})
	return m.chatService
}

// File returns the service for file management.
func (m *provider) File() ctrsvc.FileService {
	m.fileServiceOnce.Do(func() {
		m.fileService = file.NewService(m.ctx, m.sdk)
	})
	return m.fileService
}

// FineTuning returns the service for fine-tuning jobs.
func (m *provider) FineTuning() ctrsvc.FineTuningService {
	m.fineServiceOnce.Do(func() {
		m.fineTuningService = fine_tuning.NewService(m.ctx, m.sdk)
	})
	return m.fineTuningService
}

// Model returns the service for model operations.
func (m *provider) Model() ctrsvc.ModelService {
	m.modelServiceOnce.Do(func() {
		m.modelService = model.NewService(m.ctx, m.sdk)
	})
	return m.modelService
}
