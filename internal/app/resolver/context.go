package resolver

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/kylerqws/chatbot/internal/app"
)

func ResolveContext() (context.Context, context.CancelFunc) {
	if app.ModeService == ResolveMode() {
		return signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	}
	return context.WithCancel(context.Background())
}
