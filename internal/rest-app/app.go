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
	"github.com/go-seidon/local/internal/retrieving"
	"github.com/go-seidon/local/internal/serialization"

	"github.com/gorilla/mux"
)

type RestApp struct {
	config *RestAppConfig
	server app.Server
	logger logging.Logger

	healthService healthcheck.HealthCheck
}

func (a *RestApp) Run() error {
	a.logger.Infof("Running %s:%s", a.config.GetAppName(), a.config.GetAppVersion())

	err := a.healthService.Start()
	if err != nil {
		return err
	}

	a.logger.Infof("Listening on: %s", a.config.GetAddress())
	err = a.server.ListenAndServe()
	if err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (a *RestApp) Stop() error {
	a.logger.Infof("Stopping %s on: %s", a.config.GetAppName(), a.config.GetAddress())
	return a.server.Shutdown(context.Background())
}

func NewRestApp(opts ...Option) (*RestApp, error) {
	option := RestAppOption{}
	for _, opt := range opts {
		opt(&option)
	}

	if option.Config == nil {
		return nil, fmt.Errorf("invalid rest app config")
	}
	if option.Config.DBProvider != app.DB_PROVIDER_MYSQL {
		return nil, fmt.Errorf("unsupported db provider")
	}

	var logger logging.Logger
	if option.Logger != nil {
		logger = option.Logger
	} else {
		opts := []logging.Option{}

		appOpt := logging.WithAppContext(option.Config.AppName, option.Config.AppVersion)
		opts = append(opts, appOpt)

		if option.Config.AppDebug {
			debugOpt := logging.EnableDebugging()
			opts = append(opts, debugOpt)
		}

		logger = logging.NewLogrusLog(opts...)
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

	var healthService healthcheck.HealthCheck
	if option.HealthService != nil {
		healthService = option.HealthService
	} else {
		healthCheck, err := healthcheck.NewGoHealthCheck(
			healthcheck.WithLogger(logger),
			healthcheck.AddJob(inetPingJob),
			healthcheck.AddJob(appDiskJob),
		)
		if err != nil {
			return nil, err
		}
		healthService = healthCheck
	}

	var repoOpt app.RepositoryOption
	if option.Config.DBProvider == app.DB_PROVIDER_MYSQL {
		repoOpt = app.WithMySQLRepository(
			option.Config.MySQLUser, option.Config.MySQLPassword,
			option.Config.MySQLDBName, option.Config.MySQLHost,
			option.Config.MySQLPort,
		)
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

	retrieveService, err := retrieving.NewRetriever(retrieving.NewRetrieverParam{
		FileRepo:    repo.FileRepo,
		Logger:      logger,
		FileManager: fileManager,
	})
	if err != nil {
		return nil, err
	}

	serializer := serialization.NewJsonSerializer()

	router := mux.NewRouter()
	router.Use(DefaultHeaderMiddleware)
	router.HandleFunc(
		"/",
		NewRootHandler(logger, serializer, option.Config.AppName, option.Config.AppVersion),
	)
	router.HandleFunc(
		"/health",
		NewHealthCheckHandler(logger, serializer, healthService),
	).Methods(http.MethodGet)
	router.HandleFunc(
		"/file/{unique_id}",
		NewDeleteFileHandler(logger, serializer, deleteService),
	).Methods(http.MethodDelete)
	router.HandleFunc(
		"/file/{unique_id}",
		NewRetrieveFileHandler(logger, serializer, retrieveService),
	).Methods(http.MethodGet)
	router.NotFoundHandler = NewNotFoundHandler(logger, serializer)
	router.MethodNotAllowedHandler = NewMethodNotAllowedHandler(logger, serializer)

	raCfg := &RestAppConfig{
		AppName:    option.Config.AppName,
		AppVersion: option.Config.AppVersion,
		AppHost:    option.Config.RESTAppHost,
		AppPort:    option.Config.RESTAppPort,
		DbProvider: option.Config.DBProvider,
	}

	var server app.Server
	if option.Server != nil {
		server = option.Server
	} else {
		server = &http.Server{
			Addr:    raCfg.GetAddress(),
			Handler: router,
		}
	}

	app := &RestApp{
		server:        server,
		config:        raCfg,
		logger:        logger,
		healthService: healthService,
	}
	return app, nil
}
