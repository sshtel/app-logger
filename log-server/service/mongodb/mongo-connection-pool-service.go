package service

import (
	// "strconv"
	"fmt"
	"sync"
	model "github.com/sshtel/app-logger/log-server/model"
	mongodb "github.com/sshtel/app-logger/log-server/storage/mongodb"
	// utils "github.com/sshtel/app-logger/log-server/utils"
)

type Environments struct {
	MONGO_HOSTNAME string
	MONGO_PORT string
}
type MongoConnectionPoolService struct {
	wg sync.WaitGroup
	envs Environments
	serverPool map[int]*mongodb.MongoDBServer

	InputChannel chan model.LogData
}

func NewMongoConnectionPoolService (host string, port string) MongoConnectionPoolService {
	p := new (MongoConnectionPoolService)
	p.envs.MONGO_HOSTNAME = host
	p.envs.MONGO_PORT = port
	return *p
}

func (s *MongoConnectionPoolService) Terminate() {
	s.wg.Done()
}

func (s *MongoConnectionPoolService) Run(poolCount int) {
	s.serverPool = map[int]*mongodb.MongoDBServer{}

	s.InputChannel = make(chan model.LogData)

	go func() {
		s.wg.Add(1)

		// defaultHostUri := s.envs.MONGO_HOSTNAME + ":" + s.envs.MONGO_PORT

		for i := 0; i < poolCount; i++ {
			// db init
			mongoServer := mongodb.NewMongoDBServer(i, s.envs.MONGO_HOSTNAME, s.envs.MONGO_PORT)
			mongoServer.Connect()
			defer mongoServer.Disconnect(nil)
			s.serverPool[i] = mongoServer
			fmt.Println(s.serverPool[i])

		}

		// TODO
		// check s.serverPool to ensure valid connections

		go func() {
			var roundRobinPos = 0
			for {
				select {
					case v := <-s.InputChannel: {
						if (roundRobinPos >= poolCount) { roundRobinPos = 0 }
						s.serverPool[roundRobinPos].LogDataChannel <- v
						roundRobinPos++
					}
				}

			}
		}()

		fmt.Println(s.serverPool)

		s.wg.Wait()
	}()

}
