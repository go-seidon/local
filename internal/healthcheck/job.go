package healthcheck

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/InVisionApp/go-health/checkers"
	diskchk "github.com/InVisionApp/go-health/checkers/disk"
)

type NewHttpPingJobParam struct {
	Name     string
	Interval time.Duration
	Url      string
}

type NewDiskUsageJobParam struct {
	Name      string
	Interval  time.Duration
	Directory string
}

func NewHttpPingJob(p NewHttpPingJobParam) (*HealthJob, error) {
	if strings.TrimSpace(p.Name) == "" {
		return nil, fmt.Errorf("invalid name")
	}

	pingUrl, err := url.Parse(p.Url)
	if err != nil {
		return nil, err
	}

	internetConnection, err := checkers.NewHTTP(&checkers.HTTPConfig{
		URL: pingUrl,
	})
	if err != nil {
		return nil, err
	}

	job := &HealthJob{
		Name:     p.Name,
		Checker:  internetConnection,
		Interval: p.Interval,
	}
	return job, err
}

func NewDiskUsageJob(p NewDiskUsageJobParam) (*HealthJob, error) {
	if strings.TrimSpace(p.Name) == "" {
		return nil, fmt.Errorf("invalid name")
	}
	if strings.TrimSpace(p.Directory) == "" {
		return nil, fmt.Errorf("invalid directory")
	}

	appDiskChecker, err := diskchk.NewDiskUsage(&diskchk.DiskUsageConfig{
		Path:              p.Directory,
		WarningThreshold:  50,
		CriticalThreshold: 20,
	})
	if err != nil {
		return nil, err
	}

	job := &HealthJob{
		Name:     p.Name,
		Checker:  appDiskChecker,
		Interval: p.Interval,
	}
	return job, err
}
