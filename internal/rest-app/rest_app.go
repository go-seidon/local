package rest_app

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-seidon/local/internal/app"
	"github.com/go-seidon/local/internal/deleting"
	"github.com/go-seidon/local/internal/filesystem"
	"github.com/go-seidon/local/internal/healthcheck"
	"github.com/go-seidon/local/internal/logging"
	"github.com/go-seidon/local/internal/serialization"

	"github.com/gorilla/mux"
)

type restApp struct {
	server     *http.Server
	config     *RestAppConfig
	logger     logging.Logger
	serializer serialization.Serializer

	healthService healthcheck.HealthCheck
	deleteService deleting.Deleter
}

func (a *restApp) Run() error {
	a.logger.Infof("Running %s:%s", a.config.GetAppName(), a.config.GetAppVersion())

	err := a.healthService.Start()
	if err != nil {
		return err
	}

	router := mux.NewRouter()
	a.setRouter(router)

	a.server.Handler = router
	a.server.Addr = a.config.GetAddress()

	a.logger.Infof("Listening in: %s", a.config.GetAddress())
	err = a.server.ListenAndServe()
	if err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (a *restApp) Stop() error {
	return a.server.Shutdown(context.Background())
}

func (a *restApp) setRouter(router *mux.Router) {
	rootHandler := NewRootHandler(a.logger, a.serializer, a.config.GetAppName(), a.config.GetAppVersion())
	healthCheckHandler := NewHealthCheckHandler(a.logger, a.serializer, a.healthService)
	deleteFileHandler := NewDeleteFileHandler(a.logger, a.serializer, a.deleteService)

	router.Use(DefaultHeaderMiddleware)
	router.HandleFunc("/", rootHandler)
	router.HandleFunc("/health", healthCheckHandler).Methods("GET")
	router.HandleFunc("/file/{unique_id}", deleteFileHandler).Methods("DELETE")
	router.NotFoundHandler = NewNotFoundHandler(a.logger, a.serializer)
	router.MethodNotAllowedHandler = NewMethodNotAllowedHandler(a.logger, a.serializer)
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
	if opt.Config.DbProvider != app.DB_PROVIDER_MYSQL {
		return nil, fmt.Errorf("unsupported db provider")
	}

	var logger logging.Logger
	if opt.Logger != nil {
		logger = opt.Logger
	} else {
		logger = logging.NewLogrusLog(
			logging.WithAppContext(opt.Config.AppName, opt.Config.AppVersion),
		)
	}

	inetPingJob, err := healthcheck.NewHttpPingJob(healthcheck.NewHttpPingJobParam{
		Name:     "internet-connection",
		Interval: 30 * time.Second,
		Url:      "https://google.com",
	})
	if err != nil {
		return nil, err
	}

	appDiskJob, err := healthcheck.NewDiskUsageJob(healthcheck.NewDiskUsageJobParam{
		Name:      "app-disk",
		Interval:  60 * time.Second,
		Directory: "/",
	})
	if err != nil {
		return nil, err
	}

	healthService, err := healthcheck.NewGoHealthCheck(
		healthcheck.WithLogger(logger),
		healthcheck.AddJob(inetPingJob),
		healthcheck.AddJob(appDiskJob),
	)
	if err != nil {
		return nil, err
	}

	var repoOpt app.RepositoryOption
	if opt.Config.DbProvider == app.DB_PROVIDER_MYSQL {
		repoOpt = app.WithMySQLRepository("admin", "123456", "goseidon_local", "localhost", 3308)
	}
	repo, err := app.NewRepository(repoOpt)
	if err != nil {
		return nil, err
	}

	fileManager := filesystem.NewFileManager()
	deleteService, err := deleting.NewDeleter(deleting.NewDeleterParam{
		FileRepo:    repo.FileRepo,
		Logger:      logger,
		FileManager: fileManager,
	})
	if err != nil {
		return nil, err
	}

	serializer := serialization.NewJsonSerializer()

	app := &restApp{
		server:        &http.Server{},
		config:        opt.Config,
		logger:        logger,
		serializer:    serializer,
		healthService: healthService,
		deleteService: deleteService,
	}
	return app, nil
}
