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
	RetrieveFile(ctx context.Context, p RetrieveFileParam) (*RetrieveFileResult, error)
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

type RetrieveFileParam struct {
	UniqueId string
}

type RetrieveFileResult struct {
	UniqueId  string
	Name      string
	Path      string
	MimeType  string
	Extension string
	DeletedAt *int64
}
