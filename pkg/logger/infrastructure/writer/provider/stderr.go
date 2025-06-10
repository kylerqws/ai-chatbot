package provider

import (
	"io"
	"os"

	ctrcfg "github.com/kylerqws/chatbot/pkg/logger/contract/config"
	ctrwrt "github.com/kylerqws/chatbot/pkg/logger/contract/writer"
)

// stderrProvider implements Provider using os.Stderr.
type stderrProvider struct {
	config ctrcfg.Config
}

// NewStderrProvider creates a new stderr writer provider.
func NewStderrProvider(cfg ctrcfg.Config) ctrwrt.Provider {
	return &stderrProvider{config: cfg}
}

// Writer returns the os.Stderr writer.
func (*stderrProvider) Writer() io.Writer {
	return os.Stderr
}
