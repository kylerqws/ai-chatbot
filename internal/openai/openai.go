package openai

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/kylerqws/chatbot/pkg/openai"

	ctrint "github.com/kylerqws/chatbot/internal/openai/contract"
	ctrpkg "github.com/kylerqws/chatbot/pkg/openai/contract"
	ctrsvc "github.com/kylerqws/chatbot/pkg/openai/contract/service"
)

// manager provides access to OpenAI services.
type manager struct {
	sdk ctrpkg.OpenAI

	chatOnce sync.Once
	chat     ctrsvc.ChatService

	fileOnce sync.Once
	file     ctrsvc.FileService

	fineOnce sync.Once
	fine     ctrsvc.FineTuningService

	modelOnce sync.Once
	model     ctrsvc.ModelService
}

// New returns a new OpenAI manager.
func New(ctx context.Context) ctrint.OpenAI {
	sdk, err := openai.New(ctx)
	if err != nil {
		log.Fatal(fmt.Errorf("init OpenAI SDK: %w", err))
	}
	return &manager{sdk: sdk}
}

// ChatService returns the chat service.
func (m *manager) ChatService() ctrsvc.ChatService {
	m.chatOnce.Do(func() {
		m.chat = m.sdk.ChatService()
	})
	return m.chat
}

// FileService returns the file service.
func (m *manager) FileService() ctrsvc.FileService {
	m.fileOnce.Do(func() {
		m.file = m.sdk.FileService()
	})
	return m.file
}

// FineTuningService returns the fine-tuning service.
func (m *manager) FineTuningService() ctrsvc.FineTuningService {
	m.fineOnce.Do(func() {
		m.fine = m.sdk.FineTuningService()
	})
	return m.fine
}

// ModelService returns the model service.
func (m *manager) ModelService() ctrsvc.ModelService {
	m.modelOnce.Do(func() {
		m.model = m.sdk.ModelService()
	})
	return m.model
}
