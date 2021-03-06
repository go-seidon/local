package main

import (
	"fmt"
	"os"

	"github.com/go-seidon/local/internal/app"
	"github.com/go-seidon/local/internal/config"
	rest_app "github.com/go-seidon/local/internal/rest-app"
)

func main() {
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		appEnv = "local"
	}

	appConfig := app.Config{AppEnv: appEnv}

	cfgFileName := fmt.Sprintf("config/%s.toml", appConfig.AppEnv)
	tomlConfig, err := config.NewViperConfig(
		config.WithFileName(cfgFileName),
	)
	if err != nil {
		panic(err)
	}

	err = tomlConfig.LoadConfig()
	if err != nil {
		panic(err)
	}

	err = tomlConfig.ParseConfig(&appConfig)
	if err != nil {
		panic(err)
	}

	app, err := rest_app.NewRestApp(
		rest_app.WithConfig(appConfig),
	)
	if err != nil {
		panic(err)
	}

	err = app.Run()
	if err != nil {
		panic(err)
	}
}
