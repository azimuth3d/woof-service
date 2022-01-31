package db

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/azimuth3d/woof-service/pkg/schema"
	"go.mongodb.org/mongo-driver/bson"
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

func (m *MongodbRepository) InsertWoof(ctx context.Context, woof schema.Woof) (*mongo.InsertOneResult, error) {

	var result *mongo.InsertOneResult
	var err error

	if m.PresetCollection {
		result, err = m.Collection.InsertOne(ctx, woof)
	} else {
		result, err = m.Client.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("WOOF_COLLETION_NAME")).InsertOne(ctx, woof)
	}

	return result, err
}

func (m *MongodbRepository) ListWoof(ctx context.Context, skip uint64, take uint64) ([]schema.Woof, error) {
	var err error
	var cursor *mongo.Cursor

	if m.PresetCollection {
		cursor, err = m.Collection.Find(ctx, bson.M{})

	} else {
		cursor, err = m.Client.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("WOOF_COLLECTION_NAME")).Find(ctx, bson.M{})
	}

	if err != nil {
		return nil, err
	}

	var woofs []schema.Woof

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var woof schema.Woof
		cursor.Decode(&woof)
		woofs = append(woofs, woof)
	}

	return woofs, nil
}

func (m *MongodbRepository) Close() {

}
