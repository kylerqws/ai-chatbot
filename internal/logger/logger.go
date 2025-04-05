package logger

import (
	"context"

	"github.com/kylerqws/chatbot/pkg/logger"
	ctr "github.com/kylerqws/chatbot/pkg/logger/contract"
)

func New(ctx context.Context) (ctr.Logger, error) {
	return logger.New(ctx)
}
