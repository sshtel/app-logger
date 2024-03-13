package main

import (
	"fmt"
	"runtime"

	server "github.com/sshtel/app-logger/app-logger/server"
	kafka_service "github.com/sshtel/app-logger/app-logger/service/kafka"
	service_logger "github.com/sshtel/app-logger/app-logger/service/logger"
	storage_mongodb "github.com/sshtel/app-logger/app-logger/storage/mongodb"
	utils "github.com/sshtel/app-logger/app-logger/utils"
)

func main() {
	fmt.Printf(
		"Starting app logger %s/%s\n",
		runtime.GOOS,
		runtime.GOARCH,
	)

	apiServer := server.NewApiServer(utils.GetOsEnvWithDef("PORT", "8080"))
	apiServer.GoRoutineRun()

	topic := utils.GetOsEnvWithDef("KAFKA_TOPIC_APP_LOGGER", "app-logger")
	kafkaConsumer := kafka_service.KafkaConsumer{ // b == Student{"Bob", 0}
		Topic:         topic,
		ConsumerGroup: "logger",
		Connect: kafka_service.KafkaConnection{
			BootstrapServerUri: utils.GetOsEnvWithDef("KAFKA_BOOTSTRAP_SERVER_URI", "8080"),
			ClusterKey:         utils.GetOsEnvWithDef("KAFKA_CLUSTER_KEY", "8080"),
			ClusterSecret:      utils.GetOsEnvWithDef("KAFKA_CLUSTER_SECRET", "8080"),
		},
	}

	mongodbClient := storage_mongodb.NewMongoDBClient(
		utils.GetOsEnvWithDef("MONGODB_LOG_URI", "mongodb://"),
		utils.GetOsEnvWithDef("MONGODB_LOG_USERNAME", "user"),
		utils.GetOsEnvWithDef("MONGODB_LOG_PASSWORD", "password"),
	)

	mongodbLogger := service_logger.MongoDbLogger{
		DBClient:      mongodbClient,
		KafkaConsumer: &kafkaConsumer,
	}
	mongodbLogger.Run()

}
