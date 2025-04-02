package config

type Config interface {
	IsDebug() bool
	SetDebug(bool) error
}
