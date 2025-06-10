package openai

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/kylerqws/chatbot/internal/openai/enumset"
	"github.com/kylerqws/chatbot/pkg/openai"

	ctrint "github.com/kylerqws/chatbot/internal/openai/contract"
	ctrenm "github.com/kylerqws/chatbot/internal/openai/contract/enumset"

	ctrpkg "github.com/kylerqws/chatbot/pkg/openai/contract"
	ctrsvc "github.com/kylerqws/chatbot/pkg/openai/contract/service"
)

// manager provides access to OpenAI services and enum sets.
type manager struct {
	sdk ctrpkg.OpenAI

	chatServiceOnce sync.Once
	chatService     ctrsvc.ChatService

	fileServiceOnce sync.Once
	fileService     ctrsvc.FileService

	fineServiceOnce   sync.Once
	fineTuningService ctrsvc.FineTuningService

	modelServiceOnce sync.Once
	modelService     ctrsvc.ModelService

	enumManagerSetOnce sync.Once
	enumManagerSet     ctrenm.ManagerSet
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
	m.chatServiceOnce.Do(func() {
		m.chatService = m.sdk.ChatService()
	})
	return m.chatService
}

// FileService returns the file service.
func (m *manager) FileService() ctrsvc.FileService {
	m.fileServiceOnce.Do(func() {
		m.fileService = m.sdk.FileService()
	})
	return m.fileService
}

// FineTuningService returns the fine-tuning service.
func (m *manager) FineTuningService() ctrsvc.FineTuningService {
	m.fineServiceOnce.Do(func() {
		m.fineTuningService = m.sdk.FineTuningService()
	})
	return m.fineTuningService
}

// ModelService returns the model service.
func (m *manager) ModelService() ctrsvc.ModelService {
	m.modelServiceOnce.Do(func() {
		m.modelService = m.sdk.ModelService()
	})
	return m.modelService
}

// EnumManagerSet returns the OpenAI enum manager set.
func (m *manager) EnumManagerSet() ctrenm.ManagerSet {
	m.enumManagerSetOnce.Do(func() {
		m.enumManagerSet = enumset.NewManagerSet()
	})
	return m.enumManagerSet
}
