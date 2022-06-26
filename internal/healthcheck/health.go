package healthcheck

import (
	"time"
)

const (
	STATUS_OK      = "OK"
	STATUS_WARNING = "WARNING"
	STATUS_FAILED  = "FAILED"
)

type HealthCheck interface {
	Start() error
	Stop() error
	Check() (*CheckResult, error)
}

type CheckResult struct {
	Status string
	Items  map[string]CheckResultItem
}

type CheckResultItem struct {
	Name      string
	Status    string
	Error     string
	Fatal     bool
	Metadata  interface{}
	CheckedAt time.Time
}
