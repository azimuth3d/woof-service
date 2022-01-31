package db

import (
	"testing"

	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestInsertWoof(t *testing.T) {

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	t.Run("Insert woof work correctly", func(t *testing.T) {
		mt.Run("Insert woof", func(mt *mtest.T) {
			var mongoDB *MongodbRepository = NewMongodbRepositoryFromCollection(mt.Coll)

			expected := schema.woof{
				Body:      "",
				CreatedAt: "",
			}

		})

		// var databaseRepo DatabaseRepository
	})
}
