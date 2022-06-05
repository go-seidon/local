package config

type Config interface {
	Getter
	Setter
	Defaultable
}

type Defaultable interface {
	SetDefault(key string, value interface{})
}

type Getter interface {
	GetString(key string) string
	GetInt(key string) int
	Get(key string) interface{}
}

type Setter interface {
	Set(key string, value interface{})
}
