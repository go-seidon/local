package config

type Config interface {
	Getter
	Setter
}

type Getter interface {
	GetString(key string) (string, error)
	GetInt(key string) (int, error)
	Get(key string) (interface{}, error)
}

type Setter interface {
	Set(key string, value interface{}) error
	SetDefault(key string, value interface{}) error
}
