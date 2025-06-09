package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"github.com/kylerqws/chatbot/internal/app"
	"github.com/kylerqws/chatbot/internal/app/resolver"
	"github.com/kylerqws/chatbot/internal/cli"
)

// init loads the base `.env` file and an optional mode-specific `.env.{mode}` configuration file.
func init() {
	if err := godotenv.Load(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Warning: failed to load .env: %v\n", err)
	}

	mode := resolver.ResolveMode()
	if err := godotenv.Load(".env." + mode); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Warning: failed to load .env.%s: %v\n", mode, err)
	}
}

// main initializes the application and executes the CLI entry point.
func main() {
	if err := cli.Execute(app.New(resolver.ResolveContext())); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
