package config

import "time"

type Config interface {
	GetBaseURL() string
	SetBaseURL(string) error

	GetAPIKey() string
	SetAPIKey(string) error

	GetTimeout() time.Duration
	SetTimeout(int) error
}
