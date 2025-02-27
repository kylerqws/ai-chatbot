package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/kylerqws/go-chatgpt-vk/models"
	"github.com/spf13/cobra"
)

var trainCmd = &cobra.Command{
	Use:   "train",
	Short: "Train the AI model",
	Long:  "This command starts the fine-tuning process of the AI model.",
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := os.Getenv("OPENAI_API_KEY")
		if apiKey == "" {
			log.Fatal("API key not found. Ensure OPENAI_API_KEY is set.")
		}

		fmt.Println("Starting training...")
		models.Train()
		fmt.Println("Training completed.")
	},
}

func init() {
	rootCmd.AddCommand(trainCmd)
}
