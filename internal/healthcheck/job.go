package healthcheck

import (
	"time"
)

type HealthJob struct {
	Name     string
	Checker  Checker
	Interval time.Duration
}

type Checker interface {
	Status() (interface{}, error)
}

func NewHealthJobs() ([]*HealthJob, error) {
	inetChecker, err := NewInetConnectionChecker()
	if err != nil {
		return nil, err
	}

	appDiskChecker, err := NewDiskUsageChecker()
	if err != nil {
		return nil, err
	}

	jobs := []*HealthJob{
		{
			Name:     "internet-connection",
			Checker:  inetChecker,
			Interval: 30 * time.Second,
		},
		{
			Name:     "app-disk",
			Checker:  appDiskChecker,
			Interval: 60 * time.Second,
		},
	}

	return jobs, nil
}
