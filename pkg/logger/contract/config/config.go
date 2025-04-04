package config

type Config interface {
	GetWriter() string
	SetWriter(string) error

	IsDebug() bool
	SetDebug(bool) error
}
