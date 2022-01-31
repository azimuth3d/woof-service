package db

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/azimuth3d/woof-service/pkg/schema"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongodbRepository struct {
	Client           *mongo.Client
	Collection       *mongo.Collection
	PresetCollection bool
}

func NewMongodbRepository(uri string) *MongodbRepository {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	client, error := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	defer cancel()

	if error != nil {
		log.Fatal(error)

	}

	ctx, _ = context.WithTimeout(context.Background(), 2*time.Second)
	err := client.Ping(ctx, readpref.Primary())

	if err != nil {
		log.Fatalf("Error to connect mongodb %s \n", err.Error())
	}

	return &MongodbRepository{
		Client:           client,
		PresetCollection: false,
	}
}

func NewMongodbRepositoryFromCollection(c *mongo.Collection) *MongodbRepository {
	return &MongodbRepository{
		Collection:       c,
		PresetCollection: true,
	}
}

func (m *MongodbRepository) InsertWoof(ctx context.Context, woof schema.Woof) error {

	if m.PresetCollection {
		m.Collection.InsertOne(ctx, woof)
	} else {
		m.Client.Database(os.Getenv("DB_CONNECTION_STRING")).Collection(os.Getenv("WOOF_COLLETION_NAME")).InsertOne(ctx, woof)
	}

	return nil
}

func (m *MongodbRepository) ListWoof(ctx context.Context, skip uint64, take uint64) ([]schema.Woof, error) {
	return nil, nil
}

func (m *MongodbRepository) Close() {

}
