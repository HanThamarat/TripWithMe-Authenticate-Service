package adapter

import (
	"context"
	"errors"
	"time"

	core "github.com/HanThamarat/TripWithMe-Authenticate-Service/core/users"
	"github.com/HanThamarat/TripWithMe-Authenticate-Service/packages/middlewares"
	"github.com/HanThamarat/TripWithMe-Authenticate-Service/packages/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoUserRepository struct {
	collection *mongo.Collection
}

func NewMongoUserRepository(client *mongo.Client, dbName string, collectionName string) core.UserRepository {
	return &MongoUserRepository{
		collection: client.Database(dbName).Collection(collectionName),
	}
}

func (r *MongoUserRepository) Save(user core.User) (*core.User, error) {
	var users model.User

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Check if user already exists
	filter := bson.M{
		"$or": []bson.M{
			{"email": user.Email},
		},
	}
	count, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, errors.New("user already exists")
	}

	// Hash password
	hashedPassword, err := middlewares.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	now := time.Now();
	
	users.Password = hashedPassword
	users.Email = user.Email
	users.FirstName = user.FirstName
	users.LastName = user.LastName
	users.CreatedAt = now
	users.UpdatedAt = now

	_, err = r.collection.InsertOne(ctx, users)
	if err != nil {
		return nil, err
	}

	return &user, nil
}