package configs

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var sessionCollection *mongo.Collection = GetCollection(DB, "sessions")

func SetSessionTimeToLive(ttl int32) {

	sessionCollection.Indexes().DropAll(context.TODO())
	ttlIndex := mongo.IndexModel{
		Keys:    bson.M{"createdAt": 1},
		Options: options.Index().SetExpireAfterSeconds(ttl),
	}
	sessionCollection.Indexes().CreateOne(context.TODO(), ttlIndex)
}
