package repository

import (
	"context"
	"time"
)

type (
	DeleteFn func(ctx context.Context, p DeleteFnParam) error
)

type FileRepository interface {
	DeleteFile(ctx context.Context, p DeleteFileParam) (*DeleteFileResult, error)
}

type DeleteFileParam struct {
	UniqueId string
	DeleteFn DeleteFn
}

type DeleteFnParam struct {
	FilePath string
}

type DeleteFileResult struct {
	DeletedAt time.Time
}
