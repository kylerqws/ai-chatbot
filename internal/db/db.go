package db

import (
	"context"

	"github.com/kylerqws/chatbot/pkg/db"
	ctr "github.com/kylerqws/chatbot/pkg/db/contract"
)

func New(ctx context.Context) (ctr.DB, error) {
	return db.New(ctx)
}
