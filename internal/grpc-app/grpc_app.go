package grpc_app

type grpcApp struct {
}

func NewGrpcApp() (*grpcApp, error) {
	app := &grpcApp{}
	return app, nil
}
