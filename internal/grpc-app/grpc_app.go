package grpc_app

import (
	"errors"
	"net"
	"strconv"

	"github.com/go-seidon/local/cmd"
	"github.com/go-seidon/local/internal/grpc-app/pb"
	"github.com/go-seidon/local/internal/uploading"
	"google.golang.org/grpc"
)

type grpcApp struct {
	config   cmd.Config
	fileRepo uploading.FileRepository
}

func (app *grpcApp) Run() error {
	port := ":" + strconv.Itoa(app.config.GrpcPort)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		return errors.New("Failed to listen on port " + port + " because " + err.Error())
	}

	grpcServer := grpc.NewServer()
	s := NewServer(app.fileRepo, app.config)
	pb.RegisterUploadServiceServer(grpcServer, s)

	if err := grpcServer.Serve(lis); err != nil {
		return errors.New("Cannot running server: " + err.Error())
	}
	return nil
}

func NewGrpcApp(config cmd.Config, fileRepo uploading.FileRepository) (*grpcApp, error) {
	app := &grpcApp{config: config, fileRepo: fileRepo}
	return app, nil
}
