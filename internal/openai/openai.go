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

// entrypoint aggregates OpenAI service and enum providers.
type entrypoint struct {
	ctx context.Context
	sdk ctrpkg.OpenAI

	service     ctrprv.ServiceProvider
	serviceOnce sync.Once

	enum     ctrprv.EnumProvider
	enumOnce sync.Once
}

// New creates a new OpenAI entrypoint with service and enum providers.
func New(ctx context.Context) ctrint.OpenAI {
	sdk, err := openai.New(ctx)
	if err != nil {
		log.Fatal(fmt.Errorf("create OpenAI SDK: %w", err))
	}
	return &entrypoint{ctx: ctx, sdk: sdk}
}

// ServiceProvider returns the OpenAI service provider.
func (e *entrypoint) ServiceProvider() ctrprv.ServiceProvider {
	e.serviceOnce.Do(func() {
		e.service = service.NewProvider(e.ctx, e.sdk)
	})
	return e.service
}

// EnumProvider returns the OpenAI enum manager provider.
func (e *entrypoint) EnumProvider() ctrprv.EnumProvider {
	e.enumOnce.Do(func() {
		e.enum = enum.NewProvider()
	})
	return e.enum
}
