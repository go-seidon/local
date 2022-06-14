package healthcheck

import (
	"fmt"
	"time"

	"github.com/InVisionApp/go-health"
)

const (
	STATUS_OK      = "OK"
	STATUS_WARNING = "WARNING"
	STATUS_FAILED  = "FAILED"
)

type HealthService interface {
	Start() error
	Stop() error
	Check() (*CheckResult, error)
}

type CheckResultItem struct {
	Name      string
	Status    string
	Error     string
	Fatal     bool
	Metadata  interface{}
	CheckedAt time.Time
}

type CheckResult struct {
	Status string
	Items  map[string]CheckResultItem
}

func NewHealthService(jobs []*HealthJob) (*GoHealthService, error) {
	if len(jobs) == 0 {
		return nil, fmt.Errorf("invalid jobs specified")
	}

	c := health.New()

	s := &GoHealthService{
		Client: c,
		jobs:   jobs,
	}
	return s, nil
}
