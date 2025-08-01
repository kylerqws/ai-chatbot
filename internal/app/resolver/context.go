package resolver

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/kylerqws/chatbot/internal/app"
)

// ResolveContext returns a context and its cancel function.
func ResolveContext() (context.Context, context.CancelFunc) {
	if ResolveMode() == app.ModeService {
		return signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	}
	return context.WithCancel(context.Background())
}
