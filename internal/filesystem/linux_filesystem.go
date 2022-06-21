package filesystem

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type LinuxFilesystem struct{}

func (o *LinuxFilesystem) IsDirExist(dirpath string) bool {
	output := o.runShellCommand(fmt.Sprintf("ls %s", dirpath))
	return !strings.Contains(output, "No such file or directory")
}

func (o *LinuxFilesystem) WriteBinaryFileToDisk(binaryFile []byte, fullpath string) error {
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

func (o *LinuxFilesystem) runShellCommand(cmdString string) string {
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
