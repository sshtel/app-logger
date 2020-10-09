package mongodb

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	model "../../model"
)

type MongoDBServer struct {
	PoolIdx  int
	Hostname string
	Port     int
	Uri      string
	Client   *mongo.Client
	LogDataChannel (chan model.LogData)
}

func NewMongoDBServer(poolIdx int, hostname string, portStr string) *MongoDBServer {
	port, err := strconv.Atoi(portStr)
	if err != nil {
		panic(err)
	}
	return &MongoDBServer{
		PoolIdx: poolIdx,
		Hostname: hostname, Port: port,
		Uri: "mongodb://" + hostname + ":" + portStr,
		LogDataChannel: make(chan model.LogData),
	}
}

// Connect DB
func (mongodb *MongoDBServer) Connect() *mongo.Client {
	mongodb.Uri = "mongodb://" + mongodb.Hostname + ":" + strconv.Itoa(mongodb.Port)
	clientOptions := options.Client().ApplyURI(mongodb.Uri)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDBServer!")

	mongodb.Client = client


	mongodb.LogDataChannel = make(chan model.LogData)
	go func() {
		for {
			select {
			case v := <-mongodb.LogDataChannel: {
				// mongodb.Client.Insert
				fmt.Println("------------------------")
				fmt.Println(mongodb.PoolIdx)
				fmt.Println(v)
				fmt.Println("------------------------")
			}
			}
		}
	}()


	return client
}

// Disconnect DB
func (mongodb MongoDBServer) Disconnect(client *mongo.Client) {
	mongodb.Client.Disconnect(context.TODO())
	fmt.Println("Disconnected MongoDBServer")
}
