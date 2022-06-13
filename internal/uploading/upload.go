package uploading

import (
	"context"
	"time"
)

type UploadFileParam struct {
	FileData      []byte
	FileId        string
	FileName      string
	FileExtension string
	FileSize      int64
}

type UploadFileResult struct {
	FileId     string
	FileName   string
	UploadedAt time.Time
}

type Uploader interface {
	UploadFile(ctx context.Context, p UploadFileParam) (*UploadFileResult, error)
}
