package client

import "github.com/uptrace/bun"

type Client interface {
	Connect() error
	Close() error
	DB() *bun.DB
}
