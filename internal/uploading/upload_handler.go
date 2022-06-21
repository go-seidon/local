package uploading

import (
	"errors"

	"github.com/go-seidon/local/cmd"
	"github.com/go-seidon/local/internal/clock"
	"github.com/go-seidon/local/internal/filesystem"
)

func NewUploadHandler(
	fileRepo FileRepository,
	filesystem filesystem.Filesystem,
	config cmd.Config,
) (*UploadHandler, error) {
	if fileRepo == nil {
		return nil, errors.New("File repo is needed")
	}
	if filesystem == nil {
		return nil, errors.New("Filesystem is needed")
	}
	if config.FileDirectory == "" {
		return nil, errors.New("FileDirectory config is not set")
	}
	return &UploadHandler{fileRepo: fileRepo, filesystem: filesystem, config: config}, nil
}

type UploadHandler struct {
	fileRepo   FileRepository
	filesystem filesystem.Filesystem
	config     cmd.Config
}

func (o *UploadHandler) UploadFile(p UploadFileParam) (*UploadFileResult, error) {
	var err error

	file, err := NewFile(
		p.FileId,
		p.FileName,
		p.FileClientExtension,
		p.FileExtension,
		p.FileMimetype,
		o.config.FileDirectory,
		p.FileSize,
		clock.Now(),
		clock.Now(),
	)

	if err != nil {
		return nil, err
	}

	if !o.filesystem.IsDirExist(file.GetDirectoryPath()) {
		return nil, errors.New("directory is not exist")
	}

	err = o.fileRepo.Create(file)
	if err != nil {
		// deleteErr := file.DeleteBinaryFileOnDisk()
		// if deleteErr != nil {
		// 	log.Fatalf("Cannot delete file as rollback operation for %s. Please delete manually", file.GetFullpath())
		// }
		return nil, err
	}

	err = o.filesystem.WriteBinaryFileToDisk(p.FileData, file.GetFullpath())
	if err != nil {
		return nil, err
	}

	resultDTO := &UploadFileResult{
		FileId:     file.GetUniqueId(),
		FileName:   file.GetName(),
		UploadedAt: file.GetCreatedAt(),
	}
	return resultDTO, nil
}
