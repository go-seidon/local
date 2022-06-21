package uploading

type FileRepository interface {
	Create(file *File) error
	GetByUniqueId(name string) (*File, error)
}
