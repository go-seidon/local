package uploading

type FileRepository interface {
	Create(file *File) error
	GetByFilename(name string) (*File, error)
}
