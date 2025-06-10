package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"github.com/kylerqws/chatbot/internal/app"
	"github.com/kylerqws/chatbot/internal/app/resolver"
	"github.com/kylerqws/chatbot/internal/cli"
)

// init loads environment variables from the default .env file
// and an optional override file specific to the current execution mode.
func init() {
	_ = godotenv.Load()
	_ = godotenv.Load(".env." + resolver.ResolveMode())
}

// main starts the application by resolving context and executing the CLI entry point.
// Prints an error and exits with status 1 if initialization fails.
func main() {
	if err := cli.Execute(app.New(resolver.ResolveContext())); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
