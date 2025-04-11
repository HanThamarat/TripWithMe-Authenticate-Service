package adapter

import (
	"context"
	"errors"
	"time"

	authCore "github.com/HanThamarat/TripWithMe-Authenticate-Service/core/auth"
	core "github.com/HanThamarat/TripWithMe-Authenticate-Service/core/auth"
	"github.com/HanThamarat/TripWithMe-Authenticate-Service/packages/middlewares"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthMongoAdapter struct {
	collection *mongo.Collection
}

func NewAuthMongoAdapter(client *mongo.Client, dbName string, collectionName string) core.AuthRepository {
	return &AuthMongoAdapter{collection: client.Database(dbName).Collection(collectionName)}
}

func (r *AuthMongoAdapter) Authenticate(auth authCore.Auth) (authCore.AuthResponse, error) {
	var user core.MongoUser;

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.collection.FindOne(ctx, bson.M{"email": auth.Email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return authCore.AuthResponse{}, errors.New("user not found")
		}
		return authCore.AuthResponse{}, err
	}

	if !middlewares.VerifyPassword(auth.Password, user.Password) {
		return authCore.AuthResponse{}, errors.New("invalid password")
	}

	return authCore.AuthResponse{
		User: core.UserDTO{
			ID:       user.ID.Hex(),
			Password: user.Password,
			FirstName: user.FirstName,
			LastName: user.LastName,
			Email:    user.Email,
		},
	}, nil
}
