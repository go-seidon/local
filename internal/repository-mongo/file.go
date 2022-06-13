package repository_mongo

import (
	"strconv"
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/go-seidon/local/internal/uploading"
)

type FileRecord struct {
	UniqueId        string `bson:"unique_id"`
	Name            string `bson:"name"`
	Mimetype        string `bson:"mimetype"`
	Extension       string `bson:"extension"`
	ClientExtension string `bson:"client_extension"`
	Size            uint64 `bson:"size"`
	DirectoryPath   string `bson:"path"`
	CreatedAt       int64  `bson:"created_at"`
	UpdatedAt       int64  `bson:"updated_at"`
}

func NewFile(host string, port int) *File {
	session, err := mgo.Dial(host + "/" + strconv.Itoa(port))
	if err != nil {
		panic("Cannot connect to mongoDB")
	}
	return &File{
		db: session.DB("goseidon_local"),
	}
}

type File struct {
	db *mgo.Database
}

func (o *File) Create(file *uploading.File) error {
	return o.db.C("file").Insert(bson.M{
		"unique_id":        file.GetUniqueId(),
		"name":             file.GetName(),
		"mimetype":         file.GetMimetype(),
		"extension":        file.GetExtension(),
		"client_extension": file.GetClientExtension(),
		"size":             file.GetSize(),
		"path":             file.GetDirectoryPath(),
		"created_at":       file.GetCreatedAt(),
		"updated_at":       file.GetUpdatedAt(),
	})
}

func (o *File) GetByFilename(filename string) (*uploading.File, error) {
	record := FileRecord{}
	err := o.db.C("file").Find(bson.M{
		"name": filename,
	}).One(&record)

	if err != nil {
		return nil, err
	}

	file := uploading.NewFile()
	file.SetUniqueId(record.UniqueId)
	file.SetName(record.UniqueId)
	file.SetMimetype(record.Mimetype)
	file.SetExtension(record.Mimetype)
	file.SetClientExtension(record.ClientExtension)
	file.SetSize(record.Size)
	file.SetDirectoryPath(record.DirectoryPath)
	file.SetCreatedAt(time.Unix(record.CreatedAt, 0))
	file.SetUpdatedAt(time.Unix(record.UpdatedAt, 0))
	return file, nil
}

func (o *File) Close() {
	o.db.Session.Close()
}
