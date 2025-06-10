package contract

import ctrpkg "github.com/kylerqws/chatbot/pkg/logger/contract"

// Logger aggregates access to different types of loggers.
type Logger interface {
	// DB returns the logger for database logging.
	DB() ctrpkg.Logger

	// Out returns the logger for standard output.
	Out() ctrpkg.Logger

	// Err returns the logger for standard error.
	Err() ctrpkg.Logger
}
