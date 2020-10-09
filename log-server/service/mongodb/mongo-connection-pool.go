package mongo_pool_service

import (
	"fmt"
	"sync"
)

type MongoConnectionPool struct {
	wg sync.WaitGroup
	conf MongoConfig
	serverPool map[int]*MongoDBServer

	InputChannel chan MongoLogData
}

func NewMongoConnectionPool (conf MongoConfig) *MongoConnectionPool {
	p := new (MongoConnectionPool)
	p.conf = conf
	p.InputChannel = make(chan MongoLogData)

	return p
}


func (s *MongoConnectionPool) Terminate() {
	s.wg.Done()
}

func (s *MongoConnectionPool) Run() {

	poolCount := s.conf.ConnectionPoolSize
	s.serverPool = map[int]*MongoDBServer{}

	go func() {
		s.wg.Add(1)

		for i := 0; i < poolCount; i++ {
			// db init
			mongoServer := NewMongoDBServer(i, s.conf)
			mongoServer.Connect()
			defer mongoServer.Disconnect(nil)
			s.serverPool[i] = mongoServer
			fmt.Println(s.serverPool[i])

		}


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
