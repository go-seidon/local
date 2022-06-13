package main

import (
	"fmt"

	grpc_app "github.com/go-seidon/local/internal/grpc-app"
	repository_mongo "github.com/go-seidon/local/internal/repository-mongo"
)

func main() {
	fileRepo := repository_mongo.NewFile("localhost", 27017)
	defer fileRepo.Close()

	app, err := grpc_app.NewGrpcApp(fileRepo)
	if err != nil {
		fmt.Println("failed create grpc app", err)
	}
	err = app.Run()
	if err != nil {
		fmt.Println("failed run grpc app", err)
	}
}
