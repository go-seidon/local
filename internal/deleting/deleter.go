package deleting

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-seidon/local/internal/explorer"
	"github.com/go-seidon/local/internal/logging"
	"github.com/go-seidon/local/internal/repository"
)

type Deleter interface {
	DeleteFile(ctx context.Context, p DeleteFileParam) (*DeleteFileResult, error)
}

type DeleteFileParam struct {
	FileId string
}

type DeleteFileResult struct {
	DeletedAt time.Time
}

type deleter struct {
	fileRepo    repository.FileRepository
	fileManager explorer.FileManager
	log         logging.Logger
}

func (s *deleter) DeleteFile(ctx context.Context, p DeleteFileParam) (*DeleteFileResult, error) {
	s.log.Debug("In function: DeleteFile")
	defer s.log.Debug("Returning function: DeleteFile")

	if p.FileId == "" {
		return nil, fmt.Errorf("invalid file id parameter")
	}

	conn := s.fileRepo.GetConnection()
	err := conn.Start(ctx)
	if err != nil {
		return nil, err
	}

	file, err := s.fileRepo.FindFile(ctx, repository.FindFileParam{
		UniqueId:     p.FileId,
		DbConnection: conn,
	})
	if err != nil {
		connErr := conn.Rollback(ctx)
		if connErr != nil {
			return nil, connErr
		}

		if errors.Is(err, repository.ErrorRecordNotFound) {
			return nil, ErrorResourceNotFound
		}
		return nil, err
	}

	deleteRes, err := s.fileRepo.DeleteFile(ctx, repository.DeleteFileParam{
		UniqueId:     file.UniqueId,
		DbConnection: conn,
	})
	if err != nil {
		connErr := conn.Rollback(ctx)
		if connErr != nil {
			return nil, connErr
		}

		return nil, err
	}

	_, err = s.fileManager.RemoveFile(ctx, explorer.RemoveFileParam{
		Path: file.Path,
	})
	if err != nil {
		connErr := conn.Rollback(ctx)
		if connErr != nil {
			return nil, connErr
		}

		return nil, err
	}

	connErr := conn.Commit(ctx)
	if connErr != nil {
		return nil, connErr
	}

	res := &DeleteFileResult{
		DeletedAt: deleteRes.DeletedAt,
	}
	return res, nil
}

type NewDeleterParam struct {
	FileRepo    repository.FileRepository
	FileManager explorer.FileManager
	Logger      logging.Logger
}

func NewDeleter(p NewDeleterParam) (*deleter, error) {
	if p.FileRepo == nil {
		return nil, fmt.Errorf("file repo is not specified")
	}
	if p.FileManager == nil {
		return nil, fmt.Errorf("file manager is not specified")
	}
	if p.Logger == nil {
		return nil, fmt.Errorf("logger is not specified")
	}

	s := &deleter{
		fileRepo:    p.FileRepo,
		fileManager: p.FileManager,
		log:         p.Logger,
	}
	return s, nil
}
