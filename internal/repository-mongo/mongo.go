package repository_mongo

import (
	"context"
	"fmt"
	"time"

	"github.com/go-seidon/local/cmd"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoClient(config cmd.Config) (*mongo.Client, context.Context) {
	if config.MongoHost == "" {
		panic("Mongo host is not set")
	}
	if config.MongoPort == 0 {
		panic("Mongo port is not set")
	}

	url := fmt.Sprintf("mongodb://%s:%d", config.MongoHost, config.MongoPort)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.MongoConnectTimeout)*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		panic("Cannot connect to mongoDB with url " + url + ": " + err.Error())
	}

	// we need to send ping manually because mongo.Connect() doesnt returns error although she
	// cannot connect to database
	ctx, cancel = context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err = client.Ping(ctx, nil)
	if err != nil {
		panic("Cannot connect to mongoDB with url " + url + ": " + err.Error())
	}

	return client, ctx
}
