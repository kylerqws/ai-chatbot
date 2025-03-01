package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/kylerqws/chatgpt-bot/pkg/openai"
	"github.com/spf13/cobra"
)

// jobsCmd управляет заданиями OpenAI.
var jobsCmd = &cobra.Command{
	Use:   "jobs",
	Short: "Manage OpenAI jobs",
}

// createJobCmd создаёт новое задание.
var createJobCmd = &cobra.Command{
	Use:   "create [params_json]",
	Short: "Create a new job in OpenAI",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := os.Getenv("OPENAI_API_KEY")
		if apiKey == "" {
			log.Fatal("API key is required (set OPENAI_API_KEY environment variable)")
		}

		var params map[string]interface{}
		if err := json.Unmarshal([]byte(args[0]), &params); err != nil {
			log.Fatalf("Invalid JSON parameters: %v", err)
		}

		client := openai.NewJobClient(apiKey)
		jobID, err := client.CreateJob(params)
		if err != nil {
			log.Fatalf("Failed to create job: %v", err)
		}

		// Получаем информацию о новом задании
		jobInfo, err := client.GetJobInfo(jobID)
		if err != nil {
			log.Fatalf("Failed to fetch job details: %v", err)
		}

		// Выводим JSON
		output, _ := json.MarshalIndent(jobInfo, "", "  ")
		fmt.Println(string(output))
	},
}

// infoJobCmd получает информацию о задании по его ID.
var infoJobCmd = &cobra.Command{
	Use:   "info [job_id]",
	Short: "Get job information",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := os.Getenv("OPENAI_API_KEY")
		if apiKey == "" {
			log.Fatal("API key is required (set OPENAI_API_KEY environment variable)")
		}

		client := openai.NewJobClient(apiKey)
		jobInfo, err := client.GetJobInfo(args[0])
		if err != nil {
			log.Fatalf("Failed to get job info: %v", err)
		}

		// Выводим JSON
		output, _ := json.MarshalIndent(jobInfo, "", "  ")
		fmt.Println(string(output))
	},
}

// listJobsCmd получает список всех заданий.
var listJobsCmd = &cobra.Command{
	Use:   "list",
	Short: "List all OpenAI jobs",
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := os.Getenv("OPENAI_API_KEY")
		if apiKey == "" {
			log.Fatal("API key is required (set OPENAI_API_KEY environment variable)")
		}

		client := openai.NewJobClient(apiKey)
		jobs, err := client.ListJobs()
		if err != nil {
			log.Fatalf("Failed to list jobs: %v", err)
		}

		// Выводим JSON
		output, _ := json.MarshalIndent(map[string]interface{}{"jobs": jobs}, "", "  ")
		fmt.Println(string(output))
	},
}

// cancelJobCmd отменяет задание по его ID.
var cancelJobCmd = &cobra.Command{
	Use:   "cancel [job_id]",
	Short: "Cancel a specific job",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := os.Getenv("OPENAI_API_KEY")
		if apiKey == "" {
			log.Fatal("API key is required (set OPENAI_API_KEY environment variable)")
		}

		client := openai.NewJobClient(apiKey)
		err := client.CancelJob(args[0])
		if err != nil {
			log.Fatalf("Failed to cancel job: %v", err)
		}

		// Получаем обновлённую информацию о задании
		jobInfo, err := client.GetJobInfo(args[0])
		if err != nil {
			log.Fatalf("Failed to fetch job details: %v", err)
		}

		// Выводим JSON
		output, _ := json.MarshalIndent(jobInfo, "", "  ")
		fmt.Println(string(output))
	},
}

// cancelAllJobsCmd отменяет все задания.
var cancelAllJobsCmd = &cobra.Command{
	Use:   "cancel-all",
	Short: "Cancel all OpenAI jobs",
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := os.Getenv("OPENAI_API_KEY")
		if apiKey == "" {
			log.Fatal("API key is required (set OPENAI_API_KEY environment variable)")
		}

		client := openai.NewJobClient(apiKey)
		err := client.CancelAllJobs()
		if err != nil {
			log.Fatalf("Failed to cancel all jobs: %v", err)
		}

		fmt.Println("All jobs canceled successfully")
	},
}

func init() {
	// Добавляем подкоманды в jobsCmd
	jobsCmd.AddCommand(createJobCmd)
	jobsCmd.AddCommand(infoJobCmd)
	jobsCmd.AddCommand(listJobsCmd)
	jobsCmd.AddCommand(cancelJobCmd)
	jobsCmd.AddCommand(cancelAllJobsCmd)
}
