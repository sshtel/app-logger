package mongo_pool_service

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBServer struct {
	PoolIdx        int
	Hostname       string
	Port           int
	User           string
	Password       string
	Uri            string
	Client         *mongo.Client
	LogDataChannel (chan MongoLogData)
}

func NewMongoDBServer(poolIdx int, conf MongoConfig) *MongoDBServer {

	hostname := conf.Hostname
	port := conf.Port
	portStr := strconv.Itoa(conf.Port)
	userPass := ""
	if conf.User != "" {
		userPass = conf.User
		if conf.Password != "" {
			userPass = fmt.Sprintf("%s:%s", conf.User, conf.Password)
		}
	}

	uri := ""
	if userPass == "" {
		uri = fmt.Sprintf("mongodb://%s:%s", hostname, portStr)
	} else {
		uri = fmt.Sprintf("mongodb://%s@%s:%s", userPass, hostname, portStr)
	}

	return &MongoDBServer{
		PoolIdx:  poolIdx,
		Hostname: hostname, Port: port,
		User: conf.User, Password: conf.Password,
		Uri:            uri,
		LogDataChannel: make(chan MongoLogData),
	}
}

// Connect DB
func (mongodb *MongoDBServer) Connect() *mongo.Client {
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

	mongodb.LogDataChannel = make(chan MongoLogData)
	go func() {
		for {
			select {
			case v := <-mongodb.LogDataChannel:
				{
					// fmt.Println("------------------------")
					// fmt.Println(mongodb.PoolIdx)
					// fmt.Println(v)
					// fmt.Println("------------------------")
					mongodb.InsertOne(&v)
				}
			}
		}
	}()

	return client
}

// Disconnect DB
func (mongodb *MongoDBServer) Disconnect(client *mongo.Client) {
	mongodb.Client.Disconnect(context.TODO())
	fmt.Println("Disconnected MongoDBServer")
}

func (mongodb *MongoDBServer) InsertOne(data *MongoLogData) error {
	database := data.Database
	collection := data.Collection
	target := mongodb.Client.Database(database).Collection(collection)
	_, err := target.InsertOne(context.TODO(), data.Data)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
