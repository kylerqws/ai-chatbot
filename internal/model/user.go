package model

import (
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:user"`

	EntityID  int64     `bun:",pk,autoincrement"`
	ServiceID int64     `bun:",notnull"`
	UserID    string    `bun:",notnull"`
	CreatedAt time.Time `bun:",nullzero,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,default:current_timestamp"`

	Service *Service `bun:"rel:belongs-to,join:service_id=entity_id"`
}
