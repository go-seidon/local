package healthcheck

import (
	"fmt"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/InVisionApp/go-health/checkers"
	diskchk "github.com/InVisionApp/go-health/checkers/disk"
)

type HealthJob struct {
	Name     string
	Checker  Checker
	Interval time.Duration
}

type Checker interface {
	Status() (interface{}, error)
}

type NewConnectionCheckerParam struct {
	Url string
}

type NewDiskUsageCheckerParam struct {
	Directory string
}

func NewConnectionChecker(p NewConnectionCheckerParam) (Checker, error) {
	pingUrl, err := url.Parse(p.Url)
	if err != nil {
		return nil, err
	}

	internetConnection, err := checkers.NewHTTP(&checkers.HTTPConfig{
		URL: pingUrl,
	})
	return internetConnection, err
}

func NewDiskUsageChecker(p NewDiskUsageCheckerParam) (Checker, error) {
	if strings.TrimSpace(p.Directory) == "" {
		return nil, fmt.Errorf("invalid directory")
	}

	appDiskChecker, err := diskchk.NewDiskUsage(&diskchk.DiskUsageConfig{
		Path:              p.Directory,
		WarningThreshold:  50,
		CriticalThreshold: 20,
	})
	return appDiskChecker, err
}

func NewHealthJobs() ([]*HealthJob, error) {
	inetChecker, err := NewConnectionChecker(NewConnectionCheckerParam{
		Url: "https://google.com",
	})
	if err != nil {
		return nil, err
	}

	workDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	appDiskChecker, err := NewDiskUsageChecker(NewDiskUsageCheckerParam{
		Directory: workDir,
	})
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
