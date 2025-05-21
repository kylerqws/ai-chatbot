package contract

import (
	"github.com/kylerqws/chatbot/internal/app"
	"github.com/spf13/cobra"
)

type Adapter interface {
	App() *app.App
	Command() *cobra.Command

	Use() string
	SetUse(string)

	Short() string
	SetShort(string)

	Configure() *cobra.Command
	MainConfigure() *cobra.Command
}
