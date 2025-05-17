package model

import (
	"time"

	"github.com/uptrace/bun"
)

type AiChat struct {
	bun.BaseModel `bun:"table:aichat"`

	EntityID  int64     `bun:",pk,autoincrement"`
	UserID    int64     `bun:",notnull"`
	MessageID int64     `bun:",notnull"`
	CreatedAt time.Time `bun:",nullzero,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,default:current_timestamp"`

	User    *User    `bun:"rel:belongs-to,join:user_id=entity_id"`
	Message *Message `bun:"rel:belongs-to,join:message_id=entity_id"`
}
