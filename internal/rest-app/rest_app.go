package rest_app

import (
	"fmt"
	"net/http"

	"github.com/go-seidon/local/internal/app"
	"github.com/go-seidon/local/internal/deleting"
	"github.com/go-seidon/local/internal/filesystem"
	"github.com/go-seidon/local/internal/healthcheck"
	"github.com/go-seidon/local/internal/logging"
	"github.com/go-seidon/local/internal/serialization"

	"github.com/gorilla/mux"
)

type restApp struct {
	config     *RestAppConfig
	logger     logging.Logger
	serializer serialization.Serializer
}

func (a *restApp) Run() error {
	a.logger.Info("Running %s:%s", a.config.GetAppName(), a.config.GetAppVersion())

	router := mux.NewRouter()
	err := a.setRouter(router)
	if err != nil {
		return err
	}

	a.logger.Info("Listening in: %s", a.config.GetAddress())
	return http.ListenAndServe(a.config.GetAddress(), router)
}

func (a *restApp) setRouter(router *mux.Router) error {
	healthJobs, err := healthcheck.NewHealthJobs()
	if err != nil {
		return err
	}

	healthService, err := healthcheck.NewGoHealthCheck(healthJobs)
	if err != nil {
		return err
	}

	err = healthService.Start()
	if err != nil {
		return err
	}

	fileManager := filesystem.NewFileManager()

	var repoOpt app.RepositoryOption
	if a.config.DbProvider == app.DB_PROVIDER_MYSQL {
		repoOpt = app.WithMySQLRepository("admin", "123456", "goseidon_local", "localhost", 3308)
	} else {
		return fmt.Errorf("unsupported db provider")
	}

	repo, err := app.NewRepository(repoOpt)
	if err != nil {
		return err
	}

	deleteService, err := deleting.NewDeleter(deleting.NewDeleterParam{
		FileRepo:    repo.FileRepo,
		Logger:      a.logger,
		FileManager: fileManager,
	})
	if err != nil {
		return err
	}

	rootHandler := NewRootHandler(a.logger, a.serializer, a.config.GetAppName(), a.config.GetAppVersion())
	healthCheckHandler := NewHealthCheckHandler(a.logger, a.serializer, healthService)
	deleteFileHandler := NewDeleteFileHandler(a.logger, a.serializer, deleteService)

	router.Use(DefaultHeaderMiddleware)
	router.HandleFunc("/", rootHandler)
	router.HandleFunc("/health", healthCheckHandler).Methods("GET")
	router.HandleFunc("/file/{unique_id}", deleteFileHandler).Methods("DELETE")
	router.NotFoundHandler = NewNotFoundHandler(a.logger, a.serializer)
	router.MethodNotAllowedHandler = NewMethodNotAllowedHandler(a.logger, a.serializer)

	return nil
}

type RestAppConfig struct {
	AppName    string
	AppVersion string
	AppHost    string
	AppPort    int
	DbProvider string
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

	var logger logging.Logger
	if opt.Logger != nil {
		logger = opt.Logger
	} else {
		logger = logging.NewLogrusLog(
			logging.WithAppContext(opt.Config.AppName, opt.Config.AppVersion),
		)
	}

	serializer := serialization.NewJsonSerializer()

	app := &restApp{
		config:     opt.Config,
		logger:     logger,
		serializer: serializer,
	}
	return app, nil
}
