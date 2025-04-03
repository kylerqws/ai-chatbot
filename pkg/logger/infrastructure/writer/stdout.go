package writer

import (
	"io"
	"os"

	ctrcfg "github.com/kylerqws/chatbot/pkg/logger/contract/config"
	ctrwrt "github.com/kylerqws/chatbot/pkg/logger/contract/writer"
)

type stdoutProvider struct {
	config ctrcfg.Config
}

func NewStdoutProvider(cfg ctrcfg.Config) ctrwrt.Provider {
	return &stdoutProvider{config: cfg}
}

func (stdoutProvider) Writer() io.Writer {
	return os.Stdout
}
