package mongodb_service

import (
	"fmt"
	// mongodb "github.com/sshtel/app-logger/log-server/storage/mongodb"
	model "../../model"
)

var MongoServiceRef *MongoService = nil

func InitMongoRouterService() {
	if MongoServiceRef == nil {
		MongoServiceRef = new(MongoService)
	}
}

type MongoService struct {
	defHostPool MongoConnectionPoolService

	hostTable map[string]*MongoConnectionPoolService
}


func (s *MongoService) Init() {
	
	fmt.Println("Initializing MongoService..")
	s.defHostPool = NewMongoConnectionPoolService("localhost", "27017")
	s.defHostPool.Run(3)

}


func (s *MongoService) GetInputChannel(hostnickname string) chan model.LogData {
	fmt.Println(hostnickname)
	return s.defHostPool.InputChannel
}
