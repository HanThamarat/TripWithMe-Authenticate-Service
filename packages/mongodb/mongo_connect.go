package mongodb

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/HanThamarat/TripWithMe-Authenticate-Service/packages/conf"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoDatabase struct {
	Client *mongo.Client
}

var (
	once        sync.Once
	mongoClient *mongoDatabase
)

// ConnectMongoDb initializes MongoDB only once
func ConnectMongoDb(ctx context.Context, conf *conf.Config) (*mongoDatabase, error) {
	var err error

	once.Do(func() {
		// Create a timeout context
		ctx, cancel := context.WithTimeout(ctx, 20*time.Second)
		defer cancel()
		var (
			url     = conf.DB.URL
		)
		fmt.Println(url);
		clientOpts := options.Client().ApplyURI(url)
		client, connErr := mongo.Connect(ctx, clientOpts)
		if connErr != nil {
			err = connErr
			logrus.Errorf("Failed to connect to MongoDB: %v", err)
			return
		}

		// Ping to confirm connection
		if pingErr := client.Ping(ctx, nil); pingErr != nil {
			err = pingErr
			logrus.Errorf("Failed to ping MongoDB: %v", err)
			return
		}

		logrus.Info("MongoDB connection successful")
		mongoClient = &mongoDatabase{Client: client}
	})

	return mongoClient, err
}

// GetClient returns the singleton MongoDB client
func (m *mongoDatabase) GetClient() *mongo.Client {
	return mongoClient.Client
}