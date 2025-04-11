package mongodb

import "go.mongodb.org/mongo-driver/mongo"

type MongoDatabase interface {
	GetClient() *mongo.Client
}