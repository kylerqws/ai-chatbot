package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	intapp "github.com/kylerqws/chatbot/internal/app"
	intres "github.com/kylerqws/chatbot/internal/app/resolver"
	intcli "github.com/kylerqws/chatbot/internal/cli"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	mode := intres.ResolveMode()
	if intapp.ModeService == mode || intapp.ModeUtility == mode {
		_ = godotenv.Load(".env." + mode)
	}
}

func main() {
	ctx, cancel := intres.ResolveContext()
	defer cancel()

	app, err := intapp.New(ctx, cancel)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed to create application: %v\n", err)
		os.Exit(1)
	}

	err = intcli.Execute(app)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
}
