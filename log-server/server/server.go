package server

import (
	// "fmt"
	// utils "github.com/sshtel/app-logger/log-server/utils"
	// mongodb "github.com/sshtel/app-logger/log-server/storage/mongodb"
	mongoService "github.com/sshtel/app-logger/log-server/service/mongodb"
)


type Server struct {
	apiServer ApiServer
	mongoRouter mongoService.MongoRouterService
}

func (s *Server) Run() {
	s.mongoRouter.Init()
	
	
	s.apiServer = NewApiServer(&s.mongoRouter)
	s.apiServer.Run()
}
