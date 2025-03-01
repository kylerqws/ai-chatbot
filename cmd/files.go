package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/kylerqws/chatgpt-bot/pkg/openai"
	"github.com/spf13/cobra"
)

// fileCmd представляет команду "openai:files"
var fileCmd = &cobra.Command{
	Use:   "openai:files",
	Short: "Manage files in OpenAI API",
}

// uploadCmd загружает файл в OpenAI и возвращает информацию о нем.
var uploadCmd = &cobra.Command{
	Use:   "upload [file_path] [purpose]",
	Short: "Upload a file to OpenAI",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]
		purpose := args[1]

		apiKey := os.Getenv("OPENAI_API_KEY")
		if apiKey == "" {
			log.Fatal("API key is required (set OPENAI_API_KEY environment variable)")
		}

		client := openai.NewFileClient(apiKey)

		// Загружаем файл и получаем информацию о нем
		fileInfo, err := client.UploadFile(filePath, purpose)
		if err != nil {
			log.Fatalf("Failed to upload file: %v", err)
		}

		// Выводим JSON
		output, _ := json.MarshalIndent(fileInfo, "", "  ")
		fmt.Println(string(output))
	},
}

// infoCmd получает информацию о файле по его ID.
var infoCmd = &cobra.Command{
	Use:   "info [file_id]",
	Short: "Get file information",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fileID := args[0]

		apiKey := os.Getenv("OPENAI_API_KEY")
		if apiKey == "" {
			log.Fatal("API key is required (set OPENAI_API_KEY environment variable)")
		}

		client := openai.NewFileClient(apiKey)
		fileInfo, err := client.GetFileInfo(fileID)
		if err != nil {
			log.Fatalf("Failed to get file info: %v", err)
		}

		// Выводим JSON
		output, _ := json.MarshalIndent(fileInfo, "", "  ")
		fmt.Println(string(output))
	},
}

// deleteCmd удаляет файл по его ID.
var deleteCmd = &cobra.Command{
	Use:   "delete [file_id]",
	Short: "Delete a file by ID",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fileID := args[0]

		apiKey := os.Getenv("OPENAI_API_KEY")
		if apiKey == "" {
			log.Fatal("API key is required (set OPENAI_API_KEY environment variable)")
		}

		client := openai.NewFileClient(apiKey)
		err := client.DeleteFile(fileID)
		if err != nil {
			log.Fatalf("Failed to delete file: %v", err)
		}

		fmt.Printf("File %s deleted successfully!\n", fileID)
	},
}

// deleteAllCmd удаляет все файлы из OpenAI.
var deleteAllCmd = &cobra.Command{
	Use:   "delete-all",
	Short: "Delete all files from OpenAI",
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := os.Getenv("OPENAI_API_KEY")
		if apiKey == "" {
			log.Fatal("API key is required (set OPENAI_API_KEY environment variable)")
		}

		client := openai.NewFileClient(apiKey)
		err := client.DeleteAllFiles()
		if err != nil {
			log.Fatalf("Failed to delete all files: %v", err)
		}

		fmt.Println("All files deleted successfully!")
	},
}

// listCmd получает список всех загруженных файлов.
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all uploaded files",
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := os.Getenv("OPENAI_API_KEY")
		if apiKey == "" {
			log.Fatal("API key is required (set OPENAI_API_KEY environment variable)")
		}

		client := openai.NewFileClient(apiKey)
		files, err := client.ListFiles()
		if err != nil {
			log.Fatalf("Failed to list files: %v", err)
		}

		// Выводим JSON
		output, _ := json.MarshalIndent(map[string]interface{}{"files": files}, "", "  ")
		fmt.Println(string(output))
	},
}

func init() {
	// Добавляем подкоманды в fileCmd
	fileCmd.AddCommand(uploadCmd)
	fileCmd.AddCommand(infoCmd)
	fileCmd.AddCommand(deleteCmd)
	fileCmd.AddCommand(deleteAllCmd)
	fileCmd.AddCommand(listCmd)
}
