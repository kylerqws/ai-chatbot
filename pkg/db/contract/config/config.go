package config

type Config interface {
	GetDialect() string
	SetDialect(string) error

	GetDSN() string
	SetDSN(string) error

	IsDebug() bool
	SetDebug(bool) error
}
