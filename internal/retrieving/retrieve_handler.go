package retrieving

import (
	"os"

	"github.com/go-seidon/local/internal/uploading"
)

type RetrieveHandler struct{}

func (o *RetrieveHandler) UploadFile(fileRepo uploading.FileRepository, filename string) (*FileResult, error) {
	file, err := fileRepo.GetByUniqueId(filename)
	if err != nil {
		return nil, err
	}

	binaryFile, err := o.GetBinary(file)
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

func (o *RetrieveHandler) GetBinary(file *uploading.File) ([]byte, error) {
	binaryFile, err := os.ReadFile(file.GetFullpath())
	if err != nil {
		return nil, err
	}
	return binaryFile, nil
}
