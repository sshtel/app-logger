package mongodb_service

import (
	"strconv"
	"fmt"
	"sync"
	model "../../model"
	mongodb "../../storage/mongodb"
	"../../defs"
	// utils "github.com/sshtel/app-logger/log-server/utils"
)

type MongoConnectionPoolService struct {
	wg sync.WaitGroup
	conf defs.DbConfig
	serverPool map[int]*mongodb.MongoDBServer

	InputChannel chan model.LogData
}

func NewMongoConnectionPoolService (conf defs.DbConfig) MongoConnectionPoolService {
	p := new (MongoConnectionPoolService)
	p.conf = conf
	return *p
}


// func NewMongoConnectionPoolService (host string, port string) MongoConnectionPoolService {
// 	p := new (MongoConnectionPoolService)
// 	p.envs.MONGO_HOSTNAME = host
// 	p.envs.MONGO_PORT = port
// 	return *p
// }

func (s *MongoConnectionPoolService) Terminate() {
	s.wg.Done()
}

func (s *MongoConnectionPoolService) Run() {
	poolCount := s.conf.ConnPoolSize
	s.serverPool = map[int]*mongodb.MongoDBServer{}

	s.InputChannel = make(chan model.LogData)

	go func() {
		s.wg.Add(1)

		for i := 0; i < poolCount; i++ {
			// db init
			mongoServer := mongodb.NewMongoDBServer(i, s.conf.Hostname, strconv.Itoa(s.conf.Port))
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
