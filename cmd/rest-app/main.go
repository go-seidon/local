package main

import (
	"github.com/go-seidon/local/internal/app"
	rest_app "github.com/go-seidon/local/internal/rest-app"
)

func main() {
	app, err := rest_app.NewRestApp(&rest_app.NewRestAppOption{
		Config: &rest_app.RestAppConfig{
			AppName:    "local-storage",
			AppVersion: "1.0.0",
			AppHost:    "localhost",
			AppPort:    3000,
			DbProvider: app.DB_PROVIDER_MYSQL,
		},
	})
	if err != nil {
		panic(err)
	}

	err = app.Run()
	if err != nil {
		panic(err)
	}
}
