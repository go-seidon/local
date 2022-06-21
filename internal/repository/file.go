package repository

import (
	"context"
	"time"
)

type (
	DeleteFn func(ctx context.Context, p DeleteFnParam) error
)

type FileRepository interface {
	DeleteFile(ctx context.Context, p DeleteFileParam, o DeleteFileOpt) (*DeleteFileResult, error)
}

type DeleteFileParam struct {
	UniqueId string
}

type DeleteFnParam struct {
	FilePath string
}

type DeleteFileOpt struct {
	DeleteFn DeleteFn
}

type DeleteFileResult struct {
	DeletedAt time.Time
}
