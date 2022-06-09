package grpc_app

type grpcApp struct {
}

func (app *grpcApp) Run() error {
	return nil
}

func NewGrpcApp() (*grpcApp, error) {
	app := &grpcApp{}
	return app, nil
}
