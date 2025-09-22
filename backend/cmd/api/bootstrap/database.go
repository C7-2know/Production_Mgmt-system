package bootstrap

import (
	"context"
	"time"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func NewMongoDatabase(env *Env) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	uri := env.DBUri
	clientOptions := options.Client().ApplyURI(uri)
	
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB!")
	return client
}

func CheckMongoConnection(client *mongo.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	return client.Ping(ctx, nil)
}

func CloseMongoConnection(client *mongo.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := client.Disconnect(ctx); err != nil {
		log.Fatal(err)
	}
	log.Println("Disconnected from MongoDB.")
}