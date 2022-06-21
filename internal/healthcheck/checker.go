package healthcheck

import (
	"net/url"
	"os"

	"github.com/InVisionApp/go-health/checkers"
	diskchk "github.com/InVisionApp/go-health/checkers/disk"
)

func NewInetConnectionChecker() (Checker, error) {
	pingUrl, err := url.Parse("https://google.com")
	if err != nil {
		return nil, err
	}

	internetConnection, err := checkers.NewHTTP(&checkers.HTTPConfig{
		URL: pingUrl,
	})
	return internetConnection, err
}

func NewDiskUsageChecker() (Checker, error) {
	workDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	appDiskChecker, err := diskchk.NewDiskUsage(&diskchk.DiskUsageConfig{
		Path:              workDir,
		WarningThreshold:  50,
		CriticalThreshold: 20,
	})
	return appDiskChecker, err
}
