package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/kylerqws/go-chatgpt-vk/handlers"
	"github.com/kylerqws/go-chatgpt-vk/services"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the HTTP server",
	Long:  "Starts the Go server that handles API requests.",
	Run: func(cmd *cobra.Command, args []string) {
		port := os.Getenv("PORT")
		if port == "" {
			port = "5000"
		}

		vkService := services.NewVKService(handlers.GetAIResponseHandler())
		http.HandleFunc("/vk", vkService.HandleRequest)

		fmt.Println("Server started on port", port)
		log.Fatal(http.ListenAndServe(":"+port, nil))
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
