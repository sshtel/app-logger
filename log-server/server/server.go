package server

import (
	// "fmt"
	// utils "github.com/sshtel/app-logger/log-server/utils"
	// mongodb "github.com/sshtel/app-logger/log-server/storage/mongodb"
	mongoService "../service/mongodb"
)

type Server struct {
	apiServer ApiServer
}

func (s *Server) Run() {
	mongoService.InitMongoService()

	s.apiServer = NewApiServer()
	s.apiServer.Run()
}
