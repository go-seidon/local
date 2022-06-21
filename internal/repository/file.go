package repository

import (
	"context"
	"time"
)

type (
	DeleteFn func(ctx context.Context, f DeleteFileFn) error
)

type FileRepository interface {
	DeleteFile(ctx context.Context, p DeleteFileParam, o DeleteFileOpt) (*DeleteFileResult, error)
}

type DeleteFileParam struct {
	UniqueId string
}

type DeleteFileFn interface {
	DeleteFile() (*DeleteFileFnResult, error)
}

type DeleteFileFnResult struct {
	FilePath string
}

type DeleteFileOpt struct {
	DeleteFn DeleteFn
}

type DeleteFileResult struct {
	DeletedAt time.Time
}
