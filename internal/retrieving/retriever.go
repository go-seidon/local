package retrieving

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/go-seidon/local/internal/filesystem"
	"github.com/go-seidon/local/internal/logging"
	"github.com/go-seidon/local/internal/repository"
)

type Retriever interface {
	RetrieveFile(ctx context.Context, p RetrieveFileParam) (*RetrieveFileResult, error)
}

type RetrieveFileParam struct {
	FileId string
}

type RetrieveFileResult struct {
	Data      io.ReadCloser
	UniqueId  string
	Name      string
	Path      string
	MimeType  string
	Extension string
	DeletedAt *int64
}

type retriever struct {
	fileRepo    repository.FileRepository
	fileManager filesystem.FileManager
	log         logging.Logger
}

func (s *retriever) RetrieveFile(ctx context.Context, p RetrieveFileParam) (*RetrieveFileResult, error) {
	s.log.Debug("In function: RetrieveFile")
	defer s.log.Debug("Returning function: RetrieveFile")

	if p.FileId == "" {
		return nil, fmt.Errorf("invalid file id parameter")
	}

	file, err := s.fileRepo.RetrieveFile(ctx, repository.RetrieveFileParam{
		UniqueId: p.FileId,
	})
	if err != nil {
		if errors.Is(err, repository.ErrorRecordNotFound) {
			return nil, ErrorResourceNotFound
		}
		return nil, err
	}

	oRes, err := s.fileManager.OpenFile(ctx, filesystem.OpenFileParam{
		Path: file.Path,
	})
	if err != nil {
		if errors.Is(err, filesystem.ErrorFileNotFound) {
			return nil, ErrorResourceNotFound
		}
		return nil, err
	}

	res := &RetrieveFileResult{
		Data:      oRes.File,
		UniqueId:  file.UniqueId,
		Name:      file.Name,
		Path:      file.Path,
		MimeType:  file.MimeType,
		Extension: file.Extension,
	}

	return res, nil
}

type NewRetrieverParam struct {
	FileRepo    repository.FileRepository
	FileManager filesystem.FileManager
	Logger      logging.Logger
}

func NewRetriever(p NewRetrieverParam) (*retriever, error) {
	if p.FileRepo == nil {
		return nil, fmt.Errorf("file repo is not specified")
	}
	if p.FileManager == nil {
		return nil, fmt.Errorf("file manager is not specified")
	}
	if p.Logger == nil {
		return nil, fmt.Errorf("logger is not specified")
	}

	r := &retriever{
		fileRepo:    p.FileRepo,
		fileManager: p.FileManager,
		log:         p.Logger,
	}
	return r, nil
}
