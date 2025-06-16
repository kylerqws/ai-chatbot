package resolver

import (
	"github.com/kylerqws/chatbot/internal/app"
	"os"
)

// ResolveMode returns the current execution mode.
func ResolveMode() string {
	if len(os.Args) > 1 && os.Args[1] == "server" {
		return app.ModeService
	}
	return app.DefaultMode
}
