package retrieving

import (
	"context"
	"time"
)

type FileResult struct {
	BinaryFile      []byte
	UniqueId        string
	Name            string
	Mimetype        string
	Extension       string
	ClientExtension string
	Size            uint32
	DirectoryPath   string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type Retriever interface {
	RetrieveFile(ctx context.Context, filename string) (FileResult, error)
}
