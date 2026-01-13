package database

import (
	"context"
	"time"
	"fmt"

	
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Config struct {
	URI            string
	Database       string
	Username       string
	Password       string
}

type MongoDB struct {
	Client     *mongo.Client
	Database   *mongo.Database
}

var mongoInstance *MongoDB

func ConnectMongoDB(cfg Config) (*MongoDB, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(cfg.URI)
	
	if cfg.Username != "" && cfg.Password != "" {
		credential := options.Credential{
			Username: cfg.Username,
			Password: cfg.Password,
		}
		clientOptions.SetAuth(credential)
	}

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	// Ping the database to verify connection
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %v", err)
	}

	database := client.Database(cfg.Database)
	
	mongoInstance := &MongoDB{
		Client:   client,
		Database: database,
	}

	return mongoInstance, nil
}

func GetDB() *mongo.Database {
	if mongoInstance == nil {
		panic("Database not connected.")
	}

	return mongoInstance.Database
}

func GetClient() *mongo.Client {
	if mongoInstance == nil {
		panic("Database not connected.")
	}

	return mongoInstance.Client
}

func DisconnectMongoDB() error {
	if mongoInstance == nil {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := mongoInstance.Client.Disconnect(ctx)
	if err != nil {
		return fmt.Errorf("failed to disconnect MongoDB: %v", err)
	}

	mongoInstance = nil
	return nil
}