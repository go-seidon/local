package rest_app

import (
	"fmt"
	"net/http"

	"github.com/go-seidon/local/internal/logging"
	"github.com/go-seidon/local/internal/serialization"

	"github.com/gorilla/mux"
)

type restApp struct {
	config     *RestAppConfig
	logger     logging.Logger
	serializer serialization.Serializer
}

func (app *restApp) Run() error {
	app.logger.Info("Running %s:%s", app.config.GetAppName(), app.config.GetAppVersion())

	router := mux.NewRouter()
	router.Use(DefaultHeaderMiddleware)
	router.HandleFunc("/", NewRootHandler(app.logger, app.serializer, app.config.GetAppName(), app.config.GetAppVersion()))
	router.HandleFunc("/health", NewHealthCheckHandler(app.logger, app.serializer)).Methods("GET")
	router.NotFoundHandler = NewNotFoundHandler(app.logger, app.serializer)
	router.MethodNotAllowedHandler = NewMethodNotAllowedHandler(app.logger, app.serializer)

	app.logger.Info("Listening in: %s", app.config.GetAddress())

	return http.ListenAndServe(app.config.GetAddress(), router)
}

type RestAppConfig struct {
	AppName    string
	AppVersion string
	AppHost    string
	AppPort    int
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

type NewRestAppOption struct {
	Config *RestAppConfig
	Logger logging.Logger
}

func NewRestApp(opt *NewRestAppOption) (*restApp, error) {
	if opt == nil {
		return nil, fmt.Errorf("invalid rest app option")
	}
	if opt.Config == nil {
		return nil, fmt.Errorf("invalid rest app config")
	}
	if opt.Logger == nil {
		return nil, fmt.Errorf("invalid rest app logger")
	}

	serializer := serialization.NewJsonSerializer()

	app := &restApp{
		config:     opt.Config,
		logger:     opt.Logger,
		serializer: serializer,
	}
	return app, nil
}
