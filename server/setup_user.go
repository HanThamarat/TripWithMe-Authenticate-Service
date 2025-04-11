package server

import (
	adapter "github.com/HanThamarat/TripWithMe-Authenticate-Service/adapter/users"
	core "github.com/HanThamarat/TripWithMe-Authenticate-Service/core/users"
	"github.com/HanThamarat/TripWithMe-Authenticate-Service/packages/conf"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *fiberServer) InitializeUser(api fiber.Router, conf *conf.Config, client *mongo.Client) {
	dbName := conf.DB.DBNAME
	collectionName := conf.DB.COLLECTION

	userRepo := adapter.NewMongoUserRepository(client, dbName, collectionName);
	userService := core.NewUserService(userRepo);
	userHandler := adapter.NewHttpUserHandler(userService);

	
	api.Post("/user", userHandler.CreateUser);
}