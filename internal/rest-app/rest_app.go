package rest_app

type restApp struct {
}

func (app *restApp) Run() error {
	return nil
}

func NewRestApp() (*restApp, error) {
	app := &restApp{}
	return app, nil
}
