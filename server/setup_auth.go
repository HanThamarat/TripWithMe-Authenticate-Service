package server

import (
	adapter "github.com/HanThamarat/TripWithMe-Authenticate-Service/adapter/auth"
	core "github.com/HanThamarat/TripWithMe-Authenticate-Service/core/auth"
	"github.com/HanThamarat/TripWithMe-Authenticate-Service/packages/conf"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *fiberServer) InitializeAuth(api fiber.Router, conf *conf.Config,client *mongo.Client) {
	dbName := conf.DB.DBNAME
	collectionName := conf.DB.COLLECTION

	authRepo := adapter.NewAuthMongoAdapter(client, dbName, collectionName);
	authService := core.NewAuthService(authRepo);
	authHandler := adapter.NewHttpAuthHandler(authService);

	api.Post("/auth", authHandler.Authenticate);
}