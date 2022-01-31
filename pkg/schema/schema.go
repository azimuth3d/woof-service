package schema

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Woof struct {
	ID        primitive.ObjectID  `json:"_id" bson:"_id"`
	Body      string              `json:"body" bson:"body"`
	CreatedAt primitive.Timestamp `json:"createdAt" bson:"createdAt"`
}
