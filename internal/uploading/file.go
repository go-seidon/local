package uploading

import (
	"os"
	"time"
)

func NewFile() *File {
	file := &File{}
	return file
}

type File struct {
	uniqueId        string
	name            string
	clientExtension string
	extension       string
	mimetype        string
	dir             string
	createdAt       time.Time
	updatedAt       time.Time
	binaryFile      []byte
	size            *uint64
}

func (o *File) SetUniqueId(id string) *File {
	if o.uniqueId != "" {
		panic("already set unique id")
	}
	o.uniqueId = id
	return o
}

func (o *File) SetName(name string) *File {
	o.name = name
	return o
}

/**
 * This is extension from client that could be fake
 */
func (o *File) SetClientExtension(ext string) *File {
	o.clientExtension = ext
	return o
}

func (o *File) SetExtension(ext string) *File {
	o.extension = ext
	return o
}

func (o *File) SetCreatedAt(t time.Time) *File {
	o.createdAt = t
	return o
}

func (o *File) SetUpdatedAt(t time.Time) *File {
	o.updatedAt = t
	return o
}

func (o *File) SetMimetype(m string) *File {
	o.mimetype = m
	return o
}

func (o *File) SetSize(s uint64) *File {
	o.size = &s
	return o
}

/**
 * This is full directory path e.g: /usr/local/bin
 */
func (o *File) SetDirectoryPath(dir string) *File {
	o.dir = dir
	return o
}

func (o *File) GetBinary() ([]byte, error) {
	binaryFile, err := os.ReadFile(o.GetFullpath())
	if err != nil {
		return nil, err
	}
	o.binaryFile = binaryFile
	return binaryFile, nil
}

func (o *File) GetName() string {
	if o.uniqueId == "" {
		panic("unique id is needed")
	}
	return o.uniqueId + "_" + o.name
}

func (o *File) GetClientExtension() string {
	return o.clientExtension
}

func (o *File) GetExtension() string {
	return o.extension
}

func (o *File) GetMimetype() string {
	return o.mimetype
}

func (o *File) GetUniqueId() string {
	if o.uniqueId == "" {
		panic("no unique id")
	}
	return o.uniqueId
}

func (o *File) GetSize() uint64 {
	return *o.size
}

func (o *File) GetDirectoryPath() string {
	return o.dir
}

func (o *File) GetFullpath() string {
	if o.dir == "" {
		panic("Dir need to set")
	}
	return o.dir + "/" + o.name
}

func (o *File) GetCreatedAt() time.Time {
	return o.createdAt
}

func (o *File) GetUpdatedAt() time.Time {
	return o.updatedAt
}

func (o *File) DeleteBinaryFileOnDisk() error {
	return os.Remove(o.GetFullpath())
}
