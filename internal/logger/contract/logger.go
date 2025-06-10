package contract

import ctr "github.com/kylerqws/chatbot/pkg/logger/contract"

// Logger aggregates access to different types of loggers.
type Logger interface {
	// DB returns the logger for database logging.
	DB() ctr.Logger

	// Out returns the logger for standard output.
	Out() ctr.Logger

	// Err returns the logger for standard error.
	Err() ctr.Logger
}
