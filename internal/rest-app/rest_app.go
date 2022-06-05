package rest_app

type restApp struct {
}

func NewRestApp() (*restApp, error) {
	app := &restApp{}
	return app, nil
}
