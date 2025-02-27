package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "chatgpt-bot",
	Short: "ChatGPT Bot CLI",
	Long:  "CLI for managing ChatGPT bot and AI model training",
}

// Execute запускает CLI
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Command execution failed: %v", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(trainCmd)
	rootCmd.AddCommand(serveCmd)
}
