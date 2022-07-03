package config

import (
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type viperConfig struct {
	client *viper.Viper
}

func (c *viperConfig) Get(key string) (interface{}, error) {
	if !c.client.IsSet(key) {
		return nil, ErrorConfigNotFound
	}

	res := c.client.Get(key)
	return res, nil
}

func (c *viperConfig) GetBool(key string) (bool, error) {
	if !c.client.IsSet(key) {
		return false, ErrorConfigNotFound
	}

	res := c.client.GetBool(key)
	return res, nil
}

func (c *viperConfig) GetFloat64(key string) (float64, error) {
	if !c.client.IsSet(key) {
		return 0, ErrorConfigNotFound
	}

	res := c.client.GetFloat64(key)
	return res, nil
}

func (c *viperConfig) GetInt(key string) (int, error) {
	if !c.client.IsSet(key) {
		return 0, ErrorConfigNotFound
	}

	res := c.client.GetInt(key)
	return res, nil
}

func (c *viperConfig) GetString(key string) (string, error) {
	if !c.client.IsSet(key) {
		return "", ErrorConfigNotFound
	}

	res := c.client.GetString(key)
	return res, nil
}

func (c *viperConfig) GetTime(key string) (time.Time, error) {
	if !c.client.IsSet(key) {
		return time.Time{}, ErrorConfigNotFound
	}

	res := c.client.GetTime(key)
	return res, nil
}

func (c *viperConfig) GetDuration(key string) (time.Duration, error) {
	if !c.client.IsSet(key) {
		return 0, ErrorConfigNotFound
	}

	res := c.client.GetDuration(key)
	return res, nil
}

func (c *viperConfig) Set(key string, value interface{}) error {
	c.client.Set(key, value)
	return nil
}

func (c *viperConfig) SetDefault(key string, value interface{}) error {
	c.client.SetDefault(key, value)
	return nil
}

func (c *viperConfig) IsSet(key string) (bool, error) {
	res := c.client.IsSet(key)
	return res, nil
}

func (c *viperConfig) LoadConfig() error {
	return c.client.ReadInConfig()
}

func (c *viperConfig) ParseConfig(cfg interface{}) error {
	return c.client.Unmarshal(cfg, func(dc *mapstructure.DecoderConfig) {
		dc.TagName = "env"
	})
}

func NewViperConfig(opts ...Option) (*viperConfig, error) {
	option := ConfigOption{}
	for _, opt := range opts {
		opt(&option)
	}

	client := viper.New()
	if option.FileName != "" {
		client.SetConfigFile(option.FileName)
	}
	client.AllowEmptyEnv(true)
	client.AutomaticEnv()

	c := &viperConfig{
		client: client,
	}
	return c, nil
}
