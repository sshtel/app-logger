package server

type Server struct {
	apiServer ApiServer
}

func (s *Server) Run() {
	s.apiServer = NewApiServer()
	s.apiServer.Run()
}
