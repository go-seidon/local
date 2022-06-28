package filesystem

import (
	"context"
	"io/fs"
	"os"
	"time"
)

type DirectoryManager interface {
	CreateDir(ctx context.Context, p CreateDirParam) (*CreateDirResult, error)
}

type CreateDirParam struct {
	Path       string
	Permission fs.FileMode
}

type CreateDirResult struct {
	CreatedAt time.Time
}

type directoryManager struct {
}

func (dm *directoryManager) CreateDir(ctx context.Context, p CreateDirParam) (*CreateDirResult, error) {
	err := os.MkdirAll(p.Path, p.Permission)
	if err != nil {
		return nil, err
	}

	currentTimestamp := time.Now()
	res := &CreateDirResult{
		CreatedAt: currentTimestamp,
	}
	return res, nil
}

func NewDirectoryManager() *directoryManager {
	s := &directoryManager{}
	return s
}
