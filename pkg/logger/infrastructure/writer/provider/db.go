package provider

import (
	"errors"
	"io"

	ctrcfg "github.com/kylerqws/chatbot/pkg/logger/contract/config"
	ctrwrt "github.com/kylerqws/chatbot/pkg/logger/contract/writer"
)

// dbProvider implements Provider using a database writer stub.
type dbProvider struct {
	config ctrcfg.Config
}

// dbWriter is a placeholder for database-backed io.Writer.
type dbWriter struct {
	config ctrcfg.Config
}

// NewDBProvider creates a new database writer provider.
func NewDBProvider(cfg ctrcfg.Config) ctrwrt.Provider {
	return &dbProvider{config: cfg}
}

// Writer returns a placeholder DB writer.
func (*dbProvider) Writer() io.Writer {
	return &dbWriter{}
}

// Write persists log entries to the database.
func (*dbWriter) Write(_ []byte) (int, error) {
	return 0, errors.New("db writer is not implemented")
}
