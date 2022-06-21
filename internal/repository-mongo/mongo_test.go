package repository_mongo_test

import (
	"context"
	"fmt"

	"github.com/go-seidon/local/cmd"
	. "github.com/go-seidon/local/internal/repository-mongo"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/mongo"
)

var _ = Describe("File repository", Ordered, func() {
	var mongoClient *mongo.Client

	AfterEach(func() {
		if mongoClient == nil {
			return
		}
		if err := mongoClient.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	})

	var ExpectPanicWith = func(expected interface{}, config cmd.Config) {
		Expect(func() {
			mongoClient, _ = NewMongoClient(config)
		}).Should(PanicWith(expected))
	}

	It("panic if required config is not set", func() {
		ExpectPanicWith("Mongo host is not set", cmd.Config{
			MongoHost: "",
			MongoPort: 27017,
		})

		ExpectPanicWith("Mongo port is not set", cmd.Config{
			MongoHost: "localhost",
			MongoPort: 0,
		})
	})

	It("cannot connect", func() {
		const fakePort = 9000
		errorMsgRegex := fmt.Sprintf("^Cannot connect to mongoDB with url mongodb://localhost:%d", fakePort)
		ExpectPanicWith(MatchRegexp(errorMsgRegex), cmd.Config{
			MongoHost: "localhost",
			MongoPort: fakePort,
		})
	})
})
