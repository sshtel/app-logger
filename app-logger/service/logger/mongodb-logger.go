package service_logger

import (
	"context"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	kafka_service "github.com/sshtel/app-logger/app-logger/service/kafka"
	server "github.com/sshtel/app-logger/app-logger/storage/mongodb"
	bson "go.mongodb.org/mongo-driver/bson"
	options "go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDbLogger struct {
	DBClient      *server.MongoDBClient
	KafkaConsumer *kafka_service.KafkaConsumer
}

func (s *MongoDbLogger) Run() {
	s.DBClient.Connect()

	s.KafkaConsumer.AddMessageHandler(s.MessageHandler)
	s.KafkaConsumer.Run()

}

func (s *MongoDbLogger) MessageHandler(e *kafka.Message) error {
	// fmt.Printf("%% Message on from handler!! %s:\n%s\n", e.TopicPartition, string(e.Value))
	doc, err := unmarshalBaseMessage(e.Value)
	if err != nil {
		return err
	}

	// fmt.Printf("%v\n", doc)
	filter := bson.D{{"uuid", doc.UUID}}
	update := bson.D{{"$set", doc.Payload}}
	opts := options.Update().SetUpsert(true)

	coll := s.DBClient.Client.Database(doc.Database).Collection(doc.Collection)
	coll.UpdateOne(context.TODO(), filter, update, opts)
	return nil
}
