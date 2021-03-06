package rest_app

import (
	"fmt"

	"github.com/go-seidon/local/internal/app"
	"github.com/go-seidon/local/internal/healthcheck"
	"github.com/go-seidon/local/internal/logging"
)

type RestAppConfig struct {
	AppName        string
	AppVersion     string
	AppHost        string
	AppPort        int
	UploadDir      string
	UploadFormSize int64
}

func (c *RestAppConfig) GetAppName() string {
	return c.AppName
}

func (c *RestAppConfig) GetAppVersion() string {
	return c.AppVersion
}

func (c *RestAppConfig) GetAddress() string {
	return fmt.Sprintf("%s:%d", c.AppHost, c.AppPort)
}

type RestAppOption struct {
	Config        *app.Config
	Logger        logging.Logger
	Server        app.Server
	HealthService healthcheck.HealthCheck
}

type Option func(*RestAppOption)

func WithConfig(c app.Config) Option {
	return func(rao *RestAppOption) {
		rao.Config = &c
	}
}

func WithLogger(logger logging.Logger) Option {
	return func(rao *RestAppOption) {
		rao.Logger = logger
	}
}

func WithServer(server app.Server) Option {
	return func(rao *RestAppOption) {
		rao.Server = server
	}
}

func WithService(healthService healthcheck.HealthCheck) Option {
	return func(rao *RestAppOption) {
		rao.HealthService = healthService
	}
}
