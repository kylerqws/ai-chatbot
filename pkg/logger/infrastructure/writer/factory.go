package writer

import (
	"fmt"

	ctrcfg "github.com/kylerqws/chatbot/pkg/logger/contract/config"
	ctrwrt "github.com/kylerqws/chatbot/pkg/logger/contract/writer"
	"github.com/kylerqws/chatbot/pkg/logger/infrastructure/writer/provider"
)

func NewProvider(cfg ctrcfg.Config) (ctrwrt.Provider, error) {
	wrtType := cfg.GetWriter()

	switch wrtType {
	case "stdout":
		return provider.NewStdoutProvider(cfg), nil
	case "db":
		return provider.NewDBProvider(cfg), nil
	default:
		return nil, fmt.Errorf("writer: unsupported writer type %q", wrtType)
	}
}
