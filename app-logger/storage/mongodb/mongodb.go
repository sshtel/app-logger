package storage_mongodb

import (
	"context"
	"fmt"
	"log"

	mongo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBClient struct {
	Uri      string
	Username string
	Password string
	Client   *mongo.Client
}

func NewMongoDBClient(uri, username, password string) *MongoDBClient {
	return &MongoDBClient{
		Uri:      uri,
		Username: username,
		Password: password,
	}
}

// Connect DB
func (mongodb *MongoDBClient) Connect() *mongo.Client {
	clientOptions := options.Client().ApplyURI(mongodb.Uri)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDBClient!")

	mongodb.Client = client

	return client
}

// Disconnect DB
func (mongodb MongoDBClient) Disconnect(client *mongo.Client) {
	mongodb.Client.Disconnect(context.TODO())
	fmt.Println("Disconnected MongoDBClient")
}
