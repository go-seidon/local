package uploading

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

type UploadHandler struct{}

func (o *UploadHandler) UploadFile(fileRepo FileRepository, p UploadFileParam) (*UploadFileResult, error) {
	var binaryFile []byte
	var err error

	file := NewFile()
	file.SetUniqueId(uuid.NewString())
	file.SetName(p.FileName)
	file.SetClientExtension(p.FileExtension)
	file.SetDirectoryPath("/usr/local/var/goseidon")
	file.SetCreatedAt(time.Now())
	file.SetUpdatedAt(time.Now())

	if !o.isDirExist(file.GetDirectoryPath()) {
		return nil, errors.New("directory is not exist")
	}

	err = o.writeBinaryFileToDisk(binaryFile, file.GetFullpath())
	if err != nil {
		return nil, err
	}

	file.SetSize(o.calculateSize(file.GetFullpath()))
	file.SetMimetype(o.detectMimetype(file.GetFullpath()))
	file.SetExtension(o.detectExtension(file.GetMimetype(), file.GetClientExtension()))

	err = fileRepo.Create(file)
	if err != nil {
		deleteErr := file.DeleteBinaryFileOnDisk()
		if deleteErr != nil {
			log.Fatalf("Cannot delete file as rollback operation for %s. Please delete manually", file.GetFullpath())
		}
		return nil, err
	}

	resultDTO := &UploadFileResult{
		FileId:     file.GetUniqueId(),
		FileName:   file.GetName(),
		UploadedAt: file.GetCreatedAt(),
	}
	return resultDTO, nil
}

func (o *UploadHandler) writeBinaryFileToDisk(binaryFile []byte, fullpath string) error {
	f, err := os.Create(fullpath)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(binaryFile)
	if err != nil {
		return err
	}
	f.Sync()
	return nil
}

func (o *UploadHandler) detectExtension(mimetype string, clientExtension string) string {
	var ext string
	// TODO: replace data below by parse "mimetype_extension_mapping.csv" to build mappingList
	mappingList := []MimetypeAndExtensionMapping{
		{extensions: []string{"png"}, mimetype: "image/png"},
	}

	for _, mapping := range mappingList {
		if mapping.mimetype == mimetype {

			// mapping.extensions will always have at least 1 element
			if len(mapping.extensions) == 1 {
				ext = mapping.extensions[0]
			}

			// get extension that match original if multiple extension available
			originalExtensionIsCorrect := false
			for _, e := range mapping.extensions {
				if clientExtension == ext {
					originalExtensionIsCorrect = true
					ext = e
				}
			}

			// just returns first element if original extension is fake
			if !originalExtensionIsCorrect {
				ext = mapping.extensions[0]
			}
		}
	}
	return ext
}

func (o *UploadHandler) detectMimetype(filepath string) string {
	cmdString := fmt.Sprintf("file --mime-type %s | awk '{print $2}'", filepath)
	return o.runShellCommand(cmdString)
}

func (o *UploadHandler) calculateSize(filepath string) uint64 {
	cmdString := fmt.Sprintf("wc -c < %s", filepath)
	sizeString := o.runShellCommand(cmdString)
	size, err := strconv.ParseUint(sizeString, 10, 64)
	if err != nil {
		panic("Cannot convert to uint64 from " + sizeString)
	}
	return size
}

func (o *UploadHandler) runShellCommand(cmdString string) string {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command("bash", "-c", cmdString)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		panic("Cannot run this command in golang level: " + cmdString)
	}
	if stderr.String() != "" {
		panic("Cannot run this command in cmd level: " + cmdString)
	}
	return strings.TrimSpace(stdout.String())
}

func (o *UploadHandler) isDirExist(dir string) bool {
	cmdString := fmt.Sprintf("ls %s", dir)
	output := o.runShellCommand(cmdString)
	if strings.Contains(output, "No such file or directory") {
		return false
	}
	return true
}

type MimetypeAndExtensionMapping struct {
	extensions []string
	mimetype   string
}
