package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/kylerqws/chatgpt-bot/internal/filemanager"
	"github.com/spf13/cobra"
)

var filesCmd = &cobra.Command{
	Use:   "files",
	Short: "Manage files in OpenAI",
}

var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload a file to OpenAI",
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := os.Getenv("OPENAI_API_KEY")
		if apiKey == "" {
			log.Fatal("Error: OPENAI_API_KEY is missing")
		}

		filePath, _ := cmd.Flags().GetString("path")
		purpose, _ := cmd.Flags().GetString("purpose")

		if filePath == "" || purpose == "" {
			log.Fatal("Error: --path and --purpose are required")
		}

		manager := filemanager.NewManager(apiKey)
		fileID, err := manager.UploadFile(filePath, purpose)
		if err != nil {
			log.Fatalf("Failed to upload file: %v", err)
		}

		fmt.Printf("File uploaded successfully. File ID: %s\n", fileID)
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete <file_id>",
	Short: "Delete a file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := os.Getenv("OPENAI_API_KEY")
		if apiKey == "" {
			log.Fatal("Error: OPENAI_API_KEY is missing")
		}

		fileID := args[0]
		manager := filemanager.NewManager(apiKey)

		err := manager.DeleteFile(fileID)
		if err != nil {
			log.Fatalf("Failed to delete file: %v", err)
		}

		fmt.Println("File successfully deleted")
	},
}

var deleteAllCmd = &cobra.Command{
	Use:   "delete-all",
	Short: "Delete all uploaded files",
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := os.Getenv("OPENAI_API_KEY")
		if apiKey == "" {
			log.Fatal("Error: OPENAI_API_KEY is missing")
		}

		manager := filemanager.NewManager(apiKey)

		err := manager.DeleteAllFiles()
		if err != nil {
			log.Fatalf("Failed to delete all files: %v", err)
		}

		fmt.Println("All uploaded files deleted successfully")
	},
}

var infoCmd = &cobra.Command{
	Use:   "info <file_id>",
	Short: "Retrieve file info",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := os.Getenv("OPENAI_API_KEY")
		if apiKey == "" {
			log.Fatal("Error: OPENAI_API_KEY is missing")
		}

		fileID := args[0]
		manager := filemanager.NewManager(apiKey)

		fileInfo, err := manager.GetFileInfo(fileID)
		if err != nil {
			log.Fatalf("Failed to retrieve file info: %v", err)
		}

		fileInfoJSON, _ := json.MarshalIndent(fileInfo, "", "  ")
		fmt.Println(string(fileInfoJSON))
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all uploaded files",
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := os.Getenv("OPENAI_API_KEY")
		if apiKey == "" {
			log.Fatal("Error: OPENAI_API_KEY is missing")
		}

		manager := filemanager.NewManager(apiKey)

		files, err := manager.ListFiles()
		if err != nil {
			log.Fatalf("Failed to list files: %v", err)
		}

		filesJSON, _ := json.MarshalIndent(files, "", "  ")
		fmt.Println(string(filesJSON))
	},
}

func init() {
	rootCmd.AddCommand(filesCmd)

	filesCmd.AddCommand(uploadCmd)
	filesCmd.AddCommand(deleteCmd)
	filesCmd.AddCommand(deleteAllCmd)
	filesCmd.AddCommand(infoCmd)
	filesCmd.AddCommand(listCmd)

	uploadCmd.Flags().String("path", "", "File path to upload")
	uploadCmd.Flags().String("purpose", "fine-tune", "Purpose of the file (e.g., fine-tune)")
}
