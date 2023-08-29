package server

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type ApiServer struct {
	PORT string
}

func NewApiServer(port string) ApiServer {
	p := new(ApiServer)
	p.PORT = port
	// p.PORT = utils.GetOsEnvWithDef("PORT", "8080")
	return *p
}

func (s *ApiServer) GoRoutineRun() {
	go func() {
		s.Run()
	}()
}

func (s *ApiServer) Run() {

	fmt.Println("Run api-server..")

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	v1MongoRouter := e.Group("/v1/mongo")

	e.GET("/health_check", HealthCheck)

	// v1 mongo
	v1MongoRouter.GET("/info", GetInfoMongoAll)
	v1MongoRouter.GET("/:hostnickname/info", GetInfoMongoHost)
	v1MongoRouter.GET("/:hostnickname/:database/info", GetInfoMongoDatabase)
	v1MongoRouter.GET("/:hostnickname/:database/:collection/info", GetInfoMongoCollection)
	v1MongoRouter.GET("/:hostnickname/:database/:collection/data", GetDataMongoCollection)

	// Start server
	e.Logger.Fatal(e.Start(":" + s.PORT))
}
