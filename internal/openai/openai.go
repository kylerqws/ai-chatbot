package openai

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/kylerqws/chatbot/internal/openai/enum"
	"github.com/kylerqws/chatbot/internal/openai/service"
	"github.com/kylerqws/chatbot/pkg/openai"

	ctrint "github.com/kylerqws/chatbot/internal/openai/contract"
	ctrprv "github.com/kylerqws/chatbot/internal/openai/contract/provider"
	ctrpkg "github.com/kylerqws/chatbot/pkg/openai/contract"
)

// manager aggregates OpenAI service and enum providers.
type manager struct {
	ctx context.Context
	sdk ctrpkg.OpenAI

	service     ctrprv.ServiceProvider
	serviceOnce sync.Once

	enum     ctrprv.EnumProvider
	enumOnce sync.Once
}

// New creates a new OpenAI manager.
func New(ctx context.Context) ctrint.OpenAI {
	sdk, err := openai.New(ctx)
	if err != nil {
		log.Fatal(fmt.Errorf("create OpenAI SDK: %w", err))
	}
	return &manager{ctx: ctx, sdk: sdk}
}

// ServiceProvider returns the OpenAI service provider.
func (m *manager) ServiceProvider() ctrprv.ServiceProvider {
	m.serviceOnce.Do(func() {
		m.service = service.NewProvider(m.ctx, m.sdk)
	})
	return m.service
}

// EnumProvider returns the OpenAI enum manager provider.
func (m *manager) EnumProvider() ctrprv.EnumProvider {
	m.enumOnce.Do(func() {
		m.enum = enum.NewProvider()
	})
	return m.enum
}
