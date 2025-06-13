package model

import (
	"github.com/uptrace/bun"
	"time"
)

// Message is a request-response pair exchanged with the AI.
type Message struct {
	bun.BaseModel `bun:"table:message"`

	EntityID  int64     `bun:",pk,autoincrement"`
	Request   string    `bun:",notnull"`
	Response  string    `bun:",notnull"`
	CreatedAt time.Time `bun:",nullzero,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,default:current_timestamp"`
}
