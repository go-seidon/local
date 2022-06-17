package repository

import (
	"context"
	"time"
)

type FileRepository interface {
	GetConnection() Connection
	FindFile(ctx context.Context, p FindFileParam) (*FindFileResult, error)
	DeleteFile(ctx context.Context, p DeleteFileParam) (*DeleteFileResult, error)
}

type FindFileParam struct {
	UniqueId     string
	DbConnection Connection
}

type FindFileResult struct {
	UniqueId string
	Name     string
	Path     string
}

type DeleteFileParam struct {
	UniqueId     string
	DbConnection Connection
}

type DeleteFileResult struct {
	DeletedAt time.Time
}
