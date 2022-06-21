package main

import (
	"log"

	cmd "github.com/go-seidon/local/cmd"
	grpc_app "github.com/go-seidon/local/internal/grpc-app"
	repository_mongo "github.com/go-seidon/local/internal/repository-mongo"
)

func main() {
	config := cmd.Config{
		GrpcPort:            8005,
		MongoHost:           "localhost",
		MongoPort:           27017,
		MongoConnectTimeout: 1000,
		FileDirectory:       "/Users/yudi_s/delete-me",
	}

	mongoClient, ctx := repository_mongo.NewMongoClient(config)
	fileRepo := repository_mongo.NewFile(mongoClient)
	defer func() {
		if err := mongoClient.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	app, err := grpc_app.NewGrpcApp(config, fileRepo)
	if err != nil {
		log.Fatalf("failed create grpc app %v", err)
	}
	err = app.Run()
	if err != nil {
		log.Fatalf("failed run grpc app %v", err)
	}
}
