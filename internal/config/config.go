package config

import "time"

type Config interface {
	Getter
	Setter
	Manager
}

type Getter interface {
	Get(key string) (interface{}, error)
	GetBool(key string) (bool, error)
	GetFloat64(key string) (float64, error)
	GetInt(key string) (int, error)
	GetString(key string) (string, error)
	GetTime(key string) (time.Time, error)
	GetDuration(key string) (time.Duration, error)
}

type Setter interface {
	Set(key string, value interface{}) error
	SetDefault(key string, value interface{}) error
}

type Manager interface {
	IsSet(key string) (bool, error)
	LoadConfig() error
	ParseConfig(cfg interface{}) error
}

type ConfigOption struct {
	FileName string
}

type Option func(*ConfigOption)

func WithFileName(name string) Option {
	return func(co *ConfigOption) {
		co.FileName = name
	}
}
