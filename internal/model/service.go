package model

import (
	"time"

	"github.com/uptrace/bun"
)

type Service struct {
	bun.BaseModel `bun:"table:service"`

	EntityID  int64     `bun:",pk,autoincrement"`
	Code      string    `bun:",unique,notnull"`
	Name      string    `bun:",unique,notnull"`
	CreatedAt time.Time `bun:",nullzero,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,default:current_timestamp"`
}
