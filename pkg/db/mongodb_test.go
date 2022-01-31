package db

import (
	"context"
	"testing"
	"time"

	"github.com/azimuth3d/woof-service/pkg/schema"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestInsertWoof(t *testing.T) {

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	t.Run("Insert woof work correctly", func(t *testing.T) {
		mt.Run("Insert woof", func(mt *mtest.T) {
			var mongoDB *MongodbRepository = NewMongodbRepositoryFromCollection(mt.Coll)
			loc, _ := time.LoadLocation("Asia/Bangkok")

			createTime := time.Now().In(loc).Unix()

			expected := schema.Woof{
				Body:      "Woof number 1",
				CreatedAt: primitive.Timestamp{T: uint32(createTime)},
			}

			mt.AddMockResponses(bson.D{
				{
					Key:   "ok",
					Value: 1,
				},
				{Key: "value", Value: bson.D{
					{
						Key:   "result",
						Value: "ok",
					},
				}},
			})

			result, err := mongoDB.InsertWoof(context.TODO(), expected)

			assert.Nil(t, err)
			assert.NotNil(t, result)

		})

		// var databaseRepo DatabaseRepository
	})
}

func TestListWoof(t *testing.T) {

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	t.Run("List Woof work correctly", func(t *testing.T) {
		mt.Run(" list woof", func(mt *mtest.T) {
			var mongoDB *MongodbRepository = NewMongodbRepositoryFromCollection(mt.Coll)
			loc, _ := time.LoadLocation("Asia/Bangkok")

			createTime := time.Now().In(loc).Unix()

			expected1 := schema.Woof{
				Body:      "Woof number 1",
				CreatedAt: primitive.Timestamp{T: uint32(createTime)},
			}

			expected2 := schema.Woof{
				Body:      "Woof number 2",
				CreatedAt: primitive.Timestamp{T: uint32(createTime)},
			}

			var expected []schema.Woof
			expected = append(expected, expected1)
			expected = append(expected, expected2)

			first := mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
				{Key: "body", Value: "Woof number 1"},
				{Key: "createdAt", Value: primitive.Timestamp{T: uint32(createTime)}},
			})

			second := mtest.CreateCursorResponse(1, "foo.bar", mtest.NextBatch, bson.D{
				{Key: "body", Value: "Woof number 2"},
				{Key: "createdAt", Value: primitive.Timestamp{T: uint32(createTime)}},
			})

			killCursors := mtest.CreateCursorResponse(0, "foo.bar", mtest.NextBatch)

			mt.AddMockResponses(first, second, killCursors)

			result, err := mongoDB.ListWoof(context.TODO(), 0, 2)

			assert.Nil(t, err)
			assert.Equal(t, expected, result)

		})

	})
}
