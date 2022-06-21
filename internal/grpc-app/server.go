package grpc_app

import (
	"errors"
	"io"
	"log"

	"github.com/go-seidon/local/cmd"
	"github.com/go-seidon/local/internal/filesystem"
	"github.com/go-seidon/local/internal/grpc-app/pb"
	"github.com/go-seidon/local/internal/uploading"
)

func NewServer(fileRepo uploading.FileRepository, config cmd.Config) *Server {
	return &Server{fileRepo: fileRepo, config: config}
}

type Server struct {
	fileRepo uploading.FileRepository
	config   cmd.Config
}

func (s *Server) Upload(stream pb.UploadService_UploadServer) error {
	var binaryFile []byte

	for {
		chunk, err := stream.Recv()

		if err == io.EOF {
			uploadFileParam := uploading.UploadFileParam{
				FileData:            binaryFile,
				FileId:              "FileId",
				FileName:            "Filename",
				FileExtension:       "jpg",
				FileClientExtension: "jpg",
				FileMimetype:        "image/jpeg",
				FileSize:            2500,
			}
			uploaderHandler, err := uploading.NewUploadHandler(
				s.fileRepo,
				&filesystem.LinuxFilesystem{},
				s.config,
			)
			if err != nil {
				log.Fatal(err)
				return err
			}
			_, err = uploaderHandler.UploadFile(uploadFileParam)
			if err != nil {
				log.Fatal(err)
				return err
			}

			err = stream.SendAndClose(&pb.UploadResult{
				Message: "Upload received with success",
				Code:    pb.UploadStatusCode_Ok,
			})
			return err
		}

		if err != nil {
			return errors.New("failed unexpectadely while reading chunks from stream")
		}

		binaryFile = append(binaryFile, chunk.Content...)
	}
}
