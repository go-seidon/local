package uploading

import (
	"errors"
	"time"
)

func NewFile(
	uniqueId string,
	name string,
	clientExtension string,
	extension string,
	mimetype string,
	dir string,
	size uint32,
	createdAt time.Time,
	updatedAt time.Time,
) (*File, error) {
	file := &File{
		uniqueId:        uniqueId,
		name:            name,
		clientExtension: clientExtension,
		extension:       extension,
		mimetype:        mimetype,
		dirpath:         dir,
		size:            size,
		createdAt:       createdAt,
		updatedAt:       updatedAt,
	}

	err := file.Validate()

	return file, err
}

type File struct {
	uniqueId        string
	name            string
	clientExtension string
	extension       string
	mimetype        string
	dirpath         string
	size            uint32
	createdAt       time.Time
	updatedAt       time.Time
}

func (o *File) SetUniqueId(uniqueId string) *File {
	o.uniqueId = uniqueId
	return o
}

func (o *File) SetName(name string) *File {
	o.name = name
	return o
}

func (o *File) SetExtension(ext string) *File {
	o.extension = ext
	return o
}

func (o *File) SetClientExtension(ext string) *File {
	o.clientExtension = ext
	return o
}

func (o *File) SetMimetype(mimetype string) *File {
	o.mimetype = mimetype
	return o
}

func (o *File) SetDirpath(dirpath string) *File {
	o.dirpath = dirpath
	return o
}

func (o *File) SetSize(s uint32) *File {
	o.size = s
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

func (o *File) GetName() string {
	return o.name
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

func (o *File) GetSize() uint32 {
	return o.size
}

func (o *File) GetDirectoryPath() string {
	return o.dirpath
}

func (o *File) GetFullpath() string {
	return o.dirpath + "/" + o.uniqueId
}

func (o *File) GetCreatedAt() time.Time {
	return o.createdAt
}

func (o *File) GetUpdatedAt() time.Time {
	return o.updatedAt
}

func (o *File) Validate() error {
	if o.uniqueId == "" {
		return errors.New("uniqueId is mandatory")
	}
	if o.name == "" {
		return errors.New("name is mandatory")
	}
	if o.extension == "" {
		return errors.New("extension is mandatory")
	}
	if o.mimetype == "" {
		return errors.New("mimetype is mandatory")
	}
	if o.dirpath == "" {
		return errors.New("dirpath is mandatory")
	}
	if o.dirpath == "" {
		return errors.New("dirpath is mandatory")
	}
	return nil
}

func (o *File) Clone() *File {
	return &File{
		uniqueId:        o.uniqueId,
		name:            o.name,
		clientExtension: o.clientExtension,
		extension:       o.extension,
		mimetype:        o.mimetype,
		dirpath:         o.dirpath,
		size:            o.size,
		createdAt:       o.createdAt,
		updatedAt:       o.updatedAt,
	}
}

func (o *File) ToDTO() map[string]interface{} {
	return map[string]interface{}{
		"uniqueId":        o.uniqueId,
		"name":            o.name,
		"clientExtension": o.clientExtension,
		"extension":       o.extension,
		"mimetype":        o.mimetype,
		"dirpath":         o.dirpath,
		"size":            o.size,
		"createdAt":       o.createdAt,
		"updatedAt":       o.updatedAt,
	}
}
