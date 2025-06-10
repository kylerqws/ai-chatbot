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
	_ = godotenv.Load()
	_ = godotenv.Load(".env." + resolver.ResolveMode())
}

// main initializes the application and executes the CLI entry point.
func main() {
	if err := cli.Execute(app.New(resolver.ResolveContext())); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
