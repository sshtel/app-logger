package mongodb_service

import (
	"fmt"
	"../../defs"
	// mongodb "github.com/sshtel/app-logger/log-server/storage/mongodb"
	model "../../model"
)

var MongoServiceRef *MongoService = nil

func InitMongoService() {
	if MongoServiceRef == nil {
		MongoServiceRef = new(MongoService)
		MongoServiceRef.Init()
	}
}

type MongoService struct {
	defHostPool MongoConnectionPoolService
	hostTable map[string]MongoConnectionPoolService
}


func (s *MongoService) Init() {
	fmt.Println("Initializing MongoService..")

	s.hostTable = make(map[string]MongoConnectionPoolService)

	for _, v := range defs.MongoDbConfigs {
		connPool := NewMongoConnectionPoolService(v)
		s.hostTable[v.Nickname] = connPool
		connPool.Run()
	}
	
}


func (s *MongoService) GetInputChannel(hostnickname string) chan model.LogData {
	fmt.Println(hostnickname)
	return s.defHostPool.InputChannel
}
