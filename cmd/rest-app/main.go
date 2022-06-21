package main

import (
	"github.com/go-seidon/local/internal/logging"
	rest_app "github.com/go-seidon/local/internal/rest-app"
)

func main() {
	appName := "local-storage"
	appVersion := "1.0.0"

	log, err := logging.NewLogrusLog(&logging.NewLogrusLogOption{
		AppName:    appName,
		AppVersion: appVersion,
	})
	if err != nil {
		panic(err)
	}

	app, err := rest_app.NewRestApp(&rest_app.NewRestAppOption{
		Config: &rest_app.RestAppConfig{
			AppName:    appName,
			AppVersion: appVersion,
			AppHost:    "localhost",
			AppPort:    3000,
		},
		Logger: log,
	})
	if err != nil {
		panic(err)
	}

	err = app.Run()
	if err != nil {
		panic(err)
	}
}
