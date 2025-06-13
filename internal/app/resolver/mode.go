package resolver

import (
	"github.com/kylerqws/chatbot/internal/app"
	"os"
)

// ResolveMode returns the execution mode based on CLI arguments.
func ResolveMode() string {
	if len(os.Args) > 1 && os.Args[1] == "server" {
		return app.ModeService
	}
	return app.DefaultMode
}
