package writer

import (
	"fmt"
	"github.com/kylerqws/chatbot/pkg/logger/infrastructure/writer/provider"

	ctrcfg "github.com/kylerqws/chatbot/pkg/logger/contract/config"
	ctrwrt "github.com/kylerqws/chatbot/pkg/logger/contract/writer"
)

// NewProvider returns a writer provider based on the configured writer type.
func NewProvider(cfg ctrcfg.Config) (ctrwrt.Provider, error) {
	switch wt := cfg.GetWriter(); wt {
	case ctrwrt.TypeStdout:
		return provider.NewStdoutProvider(cfg), nil
	case ctrwrt.TypeStderr:
		return provider.NewStderrProvider(cfg), nil
	case ctrwrt.TypeDB:
		return provider.NewDBProvider(cfg), nil
	default:
		return nil, fmt.Errorf("unsupported writer type: %q", wt)
	}
}
