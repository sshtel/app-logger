package service

import (
	"fmt"
	// mongodb "github.com/sshtel/app-logger/log-server/storage/mongodb"
	model "../../model"
)

type MongoRouterService struct {
	defHostPool MongoConnectionPoolService

	hostTable map[string]*MongoConnectionPoolService
}


func (s *MongoRouterService) Init() {
	
	fmt.Println("Initializing MongoRouterService..")
	s.defHostPool = NewMongoConnectionPoolService("localhost", "27017")
	s.defHostPool.Run(3)

}


func (s *MongoRouterService) GetInputChannel(hostnickname string) chan model.LogData {
	fmt.Println(hostnickname)
	return s.defHostPool.InputChannel
}
