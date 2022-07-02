package healthcheck

import (
	"fmt"

	"github.com/InVisionApp/go-health"
)

type GoHealthCheck struct {
	Client HealthClient
	jobs   []*HealthJob
}

type HealthClient interface {
	AddChecks(cfgs []*health.Config) error
	Start() error
	Stop() error
	State() (map[string]health.State, bool, error)
}

func (s *GoHealthCheck) Start() error {

	cfgs := []*health.Config{}
	for _, job := range s.jobs {
		cfgs = append(cfgs, &health.Config{
			Name:     job.Name,
			Checker:  job.Checker,
			Interval: job.Interval,
		})
	}
	err := s.Client.AddChecks(cfgs)
	if err != nil {
		return err
	}

	return s.Client.Start()
}

func (s *GoHealthCheck) Stop() error {
	return s.Client.Stop()
}

func (s *GoHealthCheck) Check() (*CheckResult, error) {
	states, isFailed, err := s.Client.State()
	if err != nil {
		return nil, err
	}

	res := &CheckResult{
		Status: STATUS_FAILED,
		Items:  make(map[string]CheckResultItem),
	}
	if isFailed {
		return res, nil
	}

	totalFailed := 0
	for key, state := range states {

		status := STATUS_OK
		if state.Status == "failed" {
			status = STATUS_FAILED
			totalFailed++
		}

		res.Items[key] = CheckResultItem{
			Name:      state.Name,
			Status:    status,
			Error:     state.Err,
			Metadata:  state.Details,
			CheckedAt: state.CheckTime.UTC(),
		}
	}

	if totalFailed == 0 {
		res.Status = STATUS_OK
	} else if totalFailed != len(states) {
		res.Status = STATUS_WARNING
	}

	return res, nil
}

func NewGoHealthCheck(opts ...Option) (*GoHealthCheck, error) {
	option := HealthCheckOption{
		Jobs: []*HealthJob{},
	}
	for _, opt := range opts {
		opt(&option)
	}
	if len(option.Jobs) == 0 {
		return nil, fmt.Errorf("invalid jobs specified")
	}
	if option.Logger == nil {
		return nil, fmt.Errorf("invalid logger specified")
	}

	c := health.New()
	c.Logger = &GoHealthLog{
		Client: option.Logger,
	}

	s := &GoHealthCheck{
		Client: c,
		jobs:   option.Jobs,
	}
	return s, nil
}
