package rest_app

import (
	"io"
	"mime/multipart"
	"net/http"
	"strings"
)

type FileInfo struct {
	Name      string
	Size      int64
	Extension string
	Mimetype  string
}

func ParseMultipartFile(file io.Reader, fh *multipart.FileHeader) (*FileInfo, error) {
	info := &FileInfo{Size: fh.Size}
	info.Name = ParseFileName(fh)
	info.Extension = ParseFileExtension(fh)

	buff := make([]byte, 512)
	_, err := file.Read(buff)
	if err != nil {
		return nil, err
	}

	info.Mimetype = http.DetectContentType(buff)

	return info, nil
}

func ParseFileName(fh *multipart.FileHeader) string {
	names := strings.Split(fh.Filename, ".")
	return names[0]
}

func ParseFileExtension(fh *multipart.FileHeader) string {
	names := strings.Split(fh.Filename, ".")
	if len(names) == 1 {
		return ""
	}
	return names[len(names)-1]
}
