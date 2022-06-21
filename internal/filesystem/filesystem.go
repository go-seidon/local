package filesystem

type Filesystem interface {
	IsDirExist(dirpath string) bool
	WriteBinaryFileToDisk(binaryFile []byte, fullpath string) error
}
