package helpers

import "go.mongodb.org/mongo-driver/bson/primitive"

func GenerateMongoID() primitive.ObjectID {
	return primitive.NewObjectID()
}
