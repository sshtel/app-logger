package service_kafka

import (
	"fmt"
	"os"

	kafka "github.com/confluentinc/confluent-kafka-go/kafka"
	// "gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type KafkaConnection struct {
	BootstrapServerUri string
	ClusterKey         string
	ClusterSecret      string
}

type MessageHandlerFunc func(*kafka.Message) error

type KafkaConsumer struct {
	Topic          string
	ConsumerGroup  string
	Connect        KafkaConnection
	MessageHandler MessageHandlerFunc
	isRun          bool
}

func (s *KafkaConsumer) GoRoutineRun() {
	s.isRun = true
	go func() {
		s.Run()
	}()
}

func (s *KafkaConsumer) AddMessageHandler(handler MessageHandlerFunc) {
	s.MessageHandler = handler
}

func (s *KafkaConsumer) Run() {
	s.isRun = true
	fmt.Printf("Starting Kafka Consumer..\n")

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": s.Connect.BootstrapServerUri,
		"group.id":          s.ConsumerGroup,
		"sasl.mechanisms":   "PLAIN",
		"security.protocol": "SASL_SSL",
		"sasl.username":     s.Connect.ClusterKey,
		"sasl.password":     s.Connect.ClusterSecret,
		"auto.offset.reset": "smallest"})

	if err != nil {
		panic(err)
	}

	consumer.SubscribeTopics([]string{s.Topic}, nil)

	msg_count := 0
	MIN_COMMIT_COUNT := 5
	for s.isRun == true {
		ev := consumer.Poll(100)
		switch e := ev.(type) {
		case kafka.AssignedPartitions:
			fmt.Fprintf(os.Stderr, "%% %v\n", e)
			consumer.Assign(e.Partitions)
		case kafka.RevokedPartitions:
			fmt.Fprintf(os.Stderr, "%% %v\n", e)
			consumer.Unassign()
		case *kafka.Message:
			msg_count += 1
			if msg_count%MIN_COMMIT_COUNT == 0 {
				consumer.Commit()
			}
			if err := s.MessageHandler(e); err != nil {
				fmt.Println(err)
			}

		case kafka.PartitionEOF:
			fmt.Printf("%% Reached %v\n", e)
		case kafka.Error:
			fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
			s.isRun = false
		case kafka.OffsetsCommitted:
			{
				// fmt.Printf("OffsetCommitted %v\n", e)
			}
		default:
			{
				// fmt.Printf("Ignored %v\n", e)
			}
		}
	}

	consumer.Close()
	fmt.Printf("KafkaConsumer Connection closed\n")
}

func (s *KafkaConsumer) Close() {
	fmt.Printf("Closing KafkaConsumer Connection..\n")
	s.isRun = false
}
