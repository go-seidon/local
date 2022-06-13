package retrieving

import (
	"github.com/go-seidon/local/internal/uploading"
)

type RetrieveHandler struct{}

func (o *RetrieveHandler) UploadFile(fileRepo uploading.FileRepository, filename string) (*FileResult, error) {
	file, err := fileRepo.GetByFilename(filename)
	if err != nil {
		return nil, err
	}

	binaryFile, err := file.GetBinary()
	if err != nil {
		return nil, err
	}

	resultDTO := &FileResult{
		BinaryFile:      binaryFile,
		UniqueId:        file.GetUniqueId(),
		Name:            file.GetName(),
		Mimetype:        file.GetMimetype(),
		Extension:       file.GetExtension(),
		ClientExtension: file.GetClientExtension(),
		Size:            file.GetSize(),
		DirectoryPath:   file.GetDirectoryPath(),
		CreatedAt:       file.GetCreatedAt(),
		UpdatedAt:       file.GetUpdatedAt(),
	}

	return resultDTO, nil
}
