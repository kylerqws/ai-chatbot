package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "chatgpt-bot",
	Short: "AI ChatBot CLI",
	Long:  "CLI for managing AI ChatBot",
}

// Execute запускает CLI
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Command execution failed: %v", err)
	}
}

func init() {
	rootCmd.AddCommand(fileCmd)
	rootCmd.AddCommand(jobsCmd)
	//rootCmd.AddCommand(trainCmd)
	//rootCmd.AddCommand(serveCmd)
}
