package mongo_pool_service

import (
	"fmt"
	"errors"
	"log"
)

type MongoService struct {
	MongoDbConfigs map[string]MongoConfig
	defHostPool MongoConnectionPool
	hostTable map[string]*MongoConnectionPool
}

func (s *MongoService) Init(conf map[string]MongoConfig) {
	fmt.Println("Initializing MongoService..")

	s.MongoDbConfigs = conf
	for k := range s.MongoDbConfigs {
		fmt.Println(s.MongoDbConfigs[k])
	}

	s.hostTable = make(map[string]*MongoConnectionPool)

	for _, v := range s.MongoDbConfigs {
		connPool := NewMongoConnectionPool(v)
		s.hostTable[v.Nickname] = connPool
		connPool.Run()
	}
	
}


func (s *MongoService) GetInputChannel(hostnickname string) (chan MongoLogData, error) {
	pool := s.hostTable[hostnickname]
	if pool == nil {
		return nil, errors.New(`could not find DB ` + hostnickname)
	}
	return pool.InputChannel, nil
}

func (s *MongoService) PutData(data *MongoLogData) error {
	channel, err := s.GetInputChannel(data.HostNickname)
	if err != nil {
		log.Println(`Failed to get channel of ` + data.HostNickname)
		return errors.New(`Failed to get channle of ` + data.HostNickname)
	}
	
	go func() {
		channel <- *data
	}()
	
	return nil
}