package healthcheck

import (
	"time"

	"github.com/go-seidon/local/internal/logging"
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

type HealthJob struct {
	Name     string
	Checker  Checker
	Interval time.Duration
}

type Checker interface {
	Status() (interface{}, error)
}

type HealthCheckOption struct {
	Jobs   []*HealthJob
	Logger logging.Logger
}

type Option func(*HealthCheckOption)

func WithLogger(logger logging.Logger) Option {
	return func(hco *HealthCheckOption) {
		hco.Logger = logger
	}
}

func AddJob(job *HealthJob) Option {
	return func(hco *HealthCheckOption) {
		hco.Jobs = append(hco.Jobs, job)
	}
}

func WithJobs(jobs []*HealthJob) Option {
	return func(hco *HealthCheckOption) {
		hco.Jobs = jobs
	}
}
