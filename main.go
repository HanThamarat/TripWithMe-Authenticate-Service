package main

import (
	"context"
	"log"

	"github.com/HanThamarat/TripWithMe-Authenticate-Service/packages/conf"
	"github.com/HanThamarat/TripWithMe-Authenticate-Service/packages/mongodb"
	"github.com/HanThamarat/TripWithMe-Authenticate-Service/server"
)

// @title Authentication
// @version 1.0
// @description Authentication | Doc by Swagger.
// @contact.name Developer Team
// @contact.url https://technexify.site
// @contact.email technexify@outlook.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /api
// @schemes http https
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Use "Bearer <API_KEY>"
func main() {
	config := conf.GetConfig();
	ctx := context.Background();

	mongoDB, err := mongodb.ConnectMongoDb(ctx, config);

	if err != nil {
		log.Fatalf("MongoDB connection error: %v", err)
	}

	server.NewFiberServer(config, mongoDB).Start();
}