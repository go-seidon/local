package repository_mongo

import (
	"context"
	"time"

	"github.com/go-seidon/local/internal/uploading"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type FileRecord struct {
	UniqueId        string    `bson:"unique_id"`
	Name            string    `bson:"name"`
	Mimetype        string    `bson:"mimetype"`
	Extension       string    `bson:"extension"`
	ClientExtension string    `bson:"client_extension"`
	Size            uint32    `bson:"size"`
	DirectoryPath   string    `bson:"path"`
	CreatedAt       time.Time `bson:"created_at"`
	UpdatedAt       time.Time `bson:"updated_at"`
}

func NewFile(session *mongo.Client) *File {
	return &File{
		db: session.Database("goseidon_local"),
	}
}

type File struct {
	db *mongo.Database
}

func (o *File) Create(file *uploading.File) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := o.db.Collection("file").InsertOne(ctx, bson.D{
		{Key: "unique_id", Value: file.GetUniqueId()},
		{Key: "name", Value: file.GetName()},
		{Key: "mimetype", Value: file.GetMimetype()},
		{Key: "extension", Value: file.GetExtension()},
		{Key: "client_extension", Value: file.GetClientExtension()},
		{Key: "size", Value: file.GetSize()},
		{Key: "path", Value: file.GetDirectoryPath()},
		{Key: "created_at", Value: file.GetCreatedAt()},
		{Key: "updated_at", Value: file.GetUpdatedAt()},
	})
	return err
}

func (o *File) GetByUniqueId(filename string) (*uploading.File, error) {
	record := FileRecord{}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := o.db.Collection("file").
		FindOne(ctx, bson.D{{Key: "unique_id", Value: filename}}).
		Decode(&record)

	if err != nil {
		return nil, err
	}

	file, err := uploading.NewFile(
		record.UniqueId,
		record.Name,
		record.ClientExtension,
		record.Extension,
		record.Mimetype,
		record.DirectoryPath,
		record.Size,
		record.CreatedAt,
		record.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return file, nil
}
