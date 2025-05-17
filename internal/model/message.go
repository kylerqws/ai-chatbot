package model

import (
	"time"

	"github.com/uptrace/bun"
)

type Message struct {
	bun.BaseModel `bun:"table:message"`

	EntityID  int64     `bun:",pk,autoincrement"`
	Request   string    `bun:",notnull"`
	Response  string    `bun:",notnull"`
	CreatedAt time.Time `bun:",nullzero,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,default:current_timestamp"`
}
