package grpc_app

import (
	"github.com/go-seidon/local/internal/uploading"
)

type grpcApp struct {
	fileRepo uploading.FileRepository
}

func (app *grpcApp) Run() error {
	return nil
}

func NewGrpcApp(fileRepo uploading.FileRepository) (*grpcApp, error) {
	app := &grpcApp{fileRepo: fileRepo}
	return app, nil
}
