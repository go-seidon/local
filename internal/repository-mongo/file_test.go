package repository_mongo_test

import (
	"context"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/go-seidon/local/cmd"
	. "github.com/go-seidon/local/internal/repository-mongo"
	"github.com/go-seidon/local/internal/uploading"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestHealthCheck(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Repository Mongo Package")
}

var _ = Describe("File repository", Ordered, func() {
	var mongoClient *mongo.Client

	BeforeAll(func() {
		mongoClient, _ = NewMongoClient(cmd.Config{
			MongoHost: "localhost",
			MongoPort: 27017,
		})
	})

	AfterAll(func() {
		_, err := mongoClient.Database("goseidon_local").Collection("file").DeleteMany(context.TODO(), bson.D{})
		if err != nil {
			panic("Cannot reset datatabase")
		}

		if err := mongoClient.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	})

	createdAt, _ := time.Parse(time.RFC3339Nano, "2022-06-16T23:48:36.858Z")
	updatedAt, _ := time.Parse(time.RFC3339Nano, "2022-06-17T23:48:36.858Z")
	sample, _ := uploading.NewFile(
		"uniqueId"+strconv.Itoa(rand.Intn(999999)),
		"filename",
		"png",
		"jpg",
		"image/jpeg",
		"/directory/path",
		1000,
		createdAt,
		updatedAt,
	)

	It("Create() saves file to DB successfully", func() {
		fileRepo := NewFile(mongoClient)
		err := fileRepo.Create(sample)
		Expect(err).To(BeNil())
	})

	It("GetByUniqueId() returns correct file", func() {
		fileRepo := NewFile(mongoClient)

		storedFile, _ := fileRepo.GetByUniqueId(sample.GetUniqueId())

		Expect(storedFile).NotTo(BeNil())
		Expect(storedFile.GetUniqueId()).To(Equal(sample.GetUniqueId()))
		Expect(storedFile.GetName()).To(Equal(sample.GetName()))
		Expect(storedFile.GetClientExtension()).To(Equal(sample.GetClientExtension()))
		Expect(storedFile.GetExtension()).To(Equal(sample.GetExtension()))
		Expect(storedFile.GetMimetype()).To(Equal(sample.GetMimetype()))
		Expect(storedFile.GetDirectoryPath()).To(Equal(sample.GetDirectoryPath()))
		Expect(storedFile.GetSize()).To(Equal(sample.GetSize()))
		Expect(storedFile.GetCreatedAt()).To(Equal(sample.GetCreatedAt()))
		Expect(storedFile.GetUpdatedAt()).To(Equal(sample.GetUpdatedAt()))
	})

	It("GetByUniqueId() returns error if the data not found", func() {
		fileRepo := NewFile(mongoClient)

		storedFile, err := fileRepo.GetByUniqueId("NOT EXIST")

		Expect(storedFile).To(BeNil())
		Expect(err).NotTo(BeNil())
		Expect(err.Error()).To(Equal("mongo: no documents in result"))
	})
})
