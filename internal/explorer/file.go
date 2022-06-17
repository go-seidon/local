package explorer

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"time"
)

var (
	ErrorFileNotFound = errors.New("file not found")
)

type FileManager interface {
	IsFileExists(ctx context.Context, p IsFileExistsParam) (bool, error)
	OpenFile(ctx context.Context, p OpenFileParam) (*OpenFileResult, error)
	RemoveFile(ctx context.Context, p RemoveFileParam) (*RemoveFileResult, error)
	CreateDir(ctx context.Context, p CreateDirParam) (*CreateDirResult, error)
	SaveFile(ctx context.Context, p SaveFileParam) (*SaveFileResult, error)
	ReadFile(ctx context.Context, p ReadFileParam) (*ReadFileResult, error)
}

type IsFileExistsParam struct {
	Path string
}

type OpenFileParam struct {
	Path string
}

type OpenFileResult struct {
	File *os.File
}

type RemoveFileParam struct {
	Path string
}

type RemoveFileResult struct {
	RemovedAt time.Time
}

type CreateDirParam struct {
	Path       string
	Permission fs.FileMode
}

type CreateDirResult struct {
	CreatedAt time.Time
}

type SaveFileParam struct {
	Name       string
	Data       []byte
	Permission fs.FileMode
}

type SaveFileResult struct {
	SavedAt time.Time
}

type ReadFileParam struct {
	File *os.File
}

type ReadFileResult struct {
	Data   []byte
	ReadAt time.Time
}

type fileManager struct {
}

func (fm *fileManager) IsFileExists(ctx context.Context, p IsFileExistsParam) (bool, error) {
	_, err := os.Stat(p.Path)
	if err == nil {
		return true, nil
	}

	notExists := errors.Is(err, os.ErrNotExist)
	if notExists {
		return false, nil
	}

	return false, err
}

func (fm *fileManager) OpenFile(ctx context.Context, p OpenFileParam) (*OpenFileResult, error) {
	file, err := os.Open(p.Path)
	if err != nil {
		return nil, err
	}

	res := &OpenFileResult{
		File: file,
	}
	return res, nil
}

func (fm *fileManager) RemoveFile(ctx context.Context, p RemoveFileParam) (*RemoveFileResult, error) {
	err := os.Remove(p.Path)
	if err != nil {
		notExists := errors.Is(err, os.ErrNotExist)
		if notExists {
			return nil, ErrorFileNotFound
		}
		return nil, err
	}

	currentTimestamp := time.Now()
	res := &RemoveFileResult{
		RemovedAt: currentTimestamp,
	}
	return res, nil
}

func (fm *fileManager) CreateDir(ctx context.Context, p CreateDirParam) (*CreateDirResult, error) {
	err := os.MkdirAll(p.Path, p.Permission)
	if err != nil {
		return nil, err
	}

	currentTimestamp := time.Now()
	res := &CreateDirResult{
		CreatedAt: currentTimestamp,
	}
	return res, nil
}

func (fm *fileManager) SaveFile(ctx context.Context, p SaveFileParam) (*SaveFileResult, error) {
	err := os.WriteFile(p.Name, p.Data, p.Permission)
	if err != nil {
		return nil, err
	}

	currentTimestamp := time.Now()
	res := &SaveFileResult{
		SavedAt: currentTimestamp,
	}
	return res, nil
}

func (fm *fileManager) ReadFile(ctx context.Context, p ReadFileParam) (*ReadFileResult, error) {
	if p.File == nil {
		return nil, fmt.Errorf("invalid file")
	}

	reader := bufio.NewReader(p.File)
	bytes, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	currentTimestamp := time.Now()
	res := &ReadFileResult{
		Data:   bytes,
		ReadAt: currentTimestamp,
	}
	return res, nil
}

func NewFileManager() (*fileManager, error) {
	s := &fileManager{}
	return s, nil
}
