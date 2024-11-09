package mongodb

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/aparnasukesh/notification-svc/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	clientInstance *mongo.Client
	clientMutex    sync.Mutex
	dbInstance     *mongo.Database
	dbName         string
)

// NewMongo creates a new MongoDB client and returns a reference to the database
func NewMongo(config config.Config) (*mongo.Database, error) {
	if clientInstance == nil {
		clientMutex.Lock()
		defer clientMutex.Unlock()

		if clientInstance == nil {
			// MongoDB connection URI with appropriate authSource
			uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/?authSource=%s&authMechanism=SCRAM-SHA-256",
				config.MONGOUSER, config.MONGOPASSWORD, config.MONGOHOST, "27018", config.MONGODBNAME)

			// Create MongoDB client
			clientOpts := options.Client().ApplyURI(uri)
			client, err := mongo.NewClient(clientOpts)
			if err != nil {
				return nil, fmt.Errorf("failed to create MongoDB client: %w", err)
			}

			// Connect to MongoDB
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			err = client.Connect(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
			}

			// Check the connection
			err = client.Ping(ctx, nil)
			if err != nil {
				return nil, fmt.Errorf("failed to ping MongoDB: %w", err)
			}

			// Set clientInstance and dbInstance
			clientInstance = client
			dbName = config.MONGODBNAME
			dbInstance = clientInstance.Database(dbName)

			log.Println("Successfully connected to MongoDB")
		}
	}

	// Return the MongoDB database instance
	return dbInstance, nil
}

// GetCollection returns the specified collection from the MongoDB database
func GetCollection(collectionName string) *mongo.Collection {
	if dbInstance == nil {
		log.Fatal("Database not initialized")
	}
	return dbInstance.Collection(collectionName)
}
