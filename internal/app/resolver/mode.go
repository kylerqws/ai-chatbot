package resolver

import (
	"os"

	"github.com/kylerqws/chatbot/internal/app"
)

func ResolveMode() string {
	if len(os.Args) > 1 && os.Args[1] == "server" {
		return app.ModeService
	}
	return app.DefaultMode
}
