package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"github.com/kylerqws/chatbot/internal/app"
	"github.com/kylerqws/chatbot/internal/app/resolver"
	"github.com/kylerqws/chatbot/internal/cli"
)

// init loads environment variables from .env and specific .env.{mode} files.
func init() {
	_ = godotenv.Load()
	_ = godotenv.Load(".env." + resolver.ResolveMode())
}

// main executes the CLI entry point and exits with code 1 on failure.
func main() {
	if err := cli.Execute(app.New(resolver.ResolveContext())); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
