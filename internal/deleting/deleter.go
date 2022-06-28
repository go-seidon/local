package deleting

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-seidon/local/internal/filesystem"
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
	fileManager filesystem.FileManager
	log         logging.Logger
}

func NewDeleteFn(fileManager filesystem.FileManager) repository.DeleteFn {
	return func(ctx context.Context, r repository.DeleteFnParam) error {
		exists, err := fileManager.IsFileExists(ctx, filesystem.IsFileExistsParam{
			Path: r.FilePath,
		})
		if err != nil {
			return err
		}

		if !exists {
			return ErrorResourceNotFound
		}

		_, err = fileManager.RemoveFile(ctx, filesystem.RemoveFileParam{
			Path: r.FilePath,
		})
		if err != nil {
			return err
		}

		return nil
	}
}

func (s *deleter) DeleteFile(ctx context.Context, p DeleteFileParam) (*DeleteFileResult, error) {
	s.log.Debug("In function: DeleteFile")
	defer s.log.Debug("Returning function: DeleteFile")

	if p.FileId == "" {
		return nil, fmt.Errorf("invalid file id parameter")
	}

	delRes, err := s.fileRepo.DeleteFile(ctx, repository.DeleteFileParam{
		UniqueId: p.FileId,
		DeleteFn: NewDeleteFn(s.fileManager),
	})

	if err != nil {
		if errors.Is(err, repository.ErrorRecordNotFound) {
			return nil, ErrorResourceNotFound
		}
		return nil, err
	}

	res := &DeleteFileResult{
		DeletedAt: delRes.DeletedAt,
	}
	return res, nil
}

type NewDeleterParam struct {
	FileRepo    repository.FileRepository
	FileManager filesystem.FileManager
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
