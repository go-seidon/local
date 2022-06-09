package main

import (
	"fmt"

	grpc_app "github.com/go-seidon/local/internal/grpc-app"
)

func main() {
	app, err := grpc_app.NewGrpcApp()
	if err != nil {
		fmt.Println("failed create grpc app", err)
	}
	err = app.Run()
	if err != nil {
		fmt.Println("failed run grpc app", err)
	}
}
