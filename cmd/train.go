package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/kylerqws/chatgpt-bot/internal/training"
	"github.com/spf13/cobra"
)

var trainCmd = &cobra.Command{
	Use:   "train",
	Short: "Manage AI training jobs",
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new AI training job",
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := os.Getenv("OPENAI_API_KEY")
		if apiKey == "" {
			log.Fatal("Error: OPENAI_API_KEY is missing in .env file")
		}

		trainer := training.NewTrainer(apiKey)

		trainingFileID, _ := cmd.Flags().GetString("file-id")
		model, _ := cmd.Flags().GetString("model")

		if trainingFileID == "" || model == "" {
			log.Fatal("Error: --file-id and --model are required")
		}

		jobID, err := trainer.CreateTrainingJob(trainingFileID, model)
		if err != nil {
			log.Fatalf("Failed to create training job: %v", err)
		}

		fmt.Printf("Training job created successfully. Job ID: %s\n", jobID)
	},
}

var statusCmd = &cobra.Command{
	Use:   "status <job_id>",
	Short: "Get the status of a training job",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := os.Getenv("OPENAI_API_KEY")
		if apiKey == "" {
			log.Fatal("Error: OPENAI_API_KEY is missing in .env file")
		}

		trainer := training.NewTrainer(apiKey)
		jobID := args[0]

		jobInfo, err := trainer.GetTrainingJobInfo(jobID)
		if err != nil {
			log.Fatalf("Failed to get job info: %v", err)
		}

		// Форматируем JSON-ответ красиво
		jobInfoJSON, _ := json.MarshalIndent(jobInfo, "", "  ")
		fmt.Println(string(jobInfoJSON))
	},
}

var cancelCmd = &cobra.Command{
	Use:   "cancel <job_id>",
	Short: "Cancel a training job",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := os.Getenv("OPENAI_API_KEY")
		if apiKey == "" {
			log.Fatal("Error: OPENAI_API_KEY is missing in .env file")
		}

		trainer := training.NewTrainer(apiKey)
		jobID := args[0]

		err := trainer.CancelTrainingJob(jobID)
		if err != nil {
			log.Fatalf("Failed to cancel job: %v", err)
		}

		fmt.Println("Training job successfully canceled")
	},
}

var cancelAllCmd = &cobra.Command{
	Use:   "cancel-all",
	Short: "Cancel all running training jobs",
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := os.Getenv("OPENAI_API_KEY")
		if apiKey == "" {
			log.Fatal("Error: OPENAI_API_KEY is missing")
		}

		trainer := training.NewTrainer(apiKey)

		err := trainer.CancelAllJobs()
		if err != nil {
			log.Fatalf("Failed to cancel all jobs: %v", err)
		}

		fmt.Println("All active training jobs canceled")
	},
}

var listJobsCmd = &cobra.Command{
	Use:   "list-jobs",
	Short: "List training jobs with filtering options",
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := os.Getenv("OPENAI_API_KEY")
		if apiKey == "" {
			log.Fatal("Error: OPENAI_API_KEY is missing")
		}

		trainer := training.NewTrainer(apiKey)

		// Получаем флаги
		active, _ := cmd.Flags().GetBool("active")
		failed, _ := cmd.Flags().GetBool("failed")
		canceled, _ := cmd.Flags().GetBool("canceled")
		sinceDays, _ := cmd.Flags().GetInt("since")

		// Определяем статус для фильтрации
		var status string
		if active {
			status = training.StatusRunning
		} else if failed {
			status = training.StatusFailed
		} else if canceled {
			status = training.StatusCanceled
		}

		// Определяем дату для фильтрации (по умолчанию 1 день назад)
		since := time.Now().AddDate(0, 0, -sinceDays)

		// Получаем список jobs
		jobs, err := trainer.ListJobs(status, since)
		if err != nil {
			log.Fatalf("Failed to list jobs: %v", err)
		}

		// Выводим JSON-результат
		jobsJSON, _ := json.MarshalIndent(jobs, "", "  ")
		fmt.Println(string(jobsJSON))
	},
}

func init() {
	trainCmd.AddCommand(createCmd)
	trainCmd.AddCommand(statusCmd)
	trainCmd.AddCommand(cancelCmd)
	trainCmd.AddCommand(cancelAllCmd)
	trainCmd.AddCommand(listJobsCmd)

	createCmd.Flags().String("file-id", "", "Training file ID")
	createCmd.Flags().String("model", "", "Model name (e.g., gpt-3.5-turbo)")
	listJobsCmd.Flags().Bool("active", false, "Show only active jobs")
	listJobsCmd.Flags().Bool("failed", false, "Show only failed jobs")
	listJobsCmd.Flags().Bool("canceled", false, "Show only canceled jobs")
	listJobsCmd.Flags().Int("since", 1, "Filter jobs created in the last N days (default: 1)")
}
