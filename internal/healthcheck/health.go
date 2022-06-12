package healthcheck

type HealthService interface {
}

type healthService struct {
}

func NewHealthService() (*healthService, error) {
	s := &healthService{}
	return s, nil
}
