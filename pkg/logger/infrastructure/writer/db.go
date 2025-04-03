package writer

import (
	"fmt"
	"io"

	ctrcfg "github.com/kylerqws/chatbot/pkg/logger/contract/config"
	ctrwrt "github.com/kylerqws/chatbot/pkg/logger/contract/writer"
)

type (
	dbProvider struct {
		config ctrcfg.Config
	}
	dbWriter struct {
		config ctrcfg.Config
	}
)

func NewDBProvider(cfg ctrcfg.Config) ctrwrt.Provider {
	return &dbProvider{config: cfg}
}

func (p *dbProvider) Writer() io.Writer {
	return &dbWriter{config: p.config}
}

// TODO: need to implement for storing logs in the DB
func (w *dbWriter) Write(p []byte) (int, error) {
	return len(p), fmt.Errorf("database writer not implemented")
}
