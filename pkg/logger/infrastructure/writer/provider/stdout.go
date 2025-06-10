package provider

import (
	"io"
	"os"

	ctrcfg "github.com/kylerqws/chatbot/pkg/logger/contract/config"
	ctrwrt "github.com/kylerqws/chatbot/pkg/logger/contract/writer"
)

// stdoutProvider implements Provider using os.Stdout.
type stdoutProvider struct {
	config ctrcfg.Config
}

// NewStdoutProvider creates a new stdout writer provider.
func NewStdoutProvider(cfg ctrcfg.Config) ctrwrt.Provider {
	return &stdoutProvider{config: cfg}
}

// Writer returns the os.Stdout writer.
func (*stdoutProvider) Writer() io.Writer {
	return os.Stdout
}
