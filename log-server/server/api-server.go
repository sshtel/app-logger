package server

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sshtel/app-logger/log-server/handlers/handler_common"
	v1 "github.com/sshtel/app-logger/log-server/handlers/handler_v1"
	utils "github.com/sshtel/app-logger/log-server/utils"
)

type ApiServer struct {
	APISERVER_PORT string
}

func NewApiServer() ApiServer {
	p := new(ApiServer)
	p.APISERVER_PORT = utils.GetOsEnvWithDef("APISERVER_PORT", "8080")
	return *p
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

	e.GET("/ping", handler_common.Pong)

	// v1 mongo
	v1MongoRouter.GET("/info", v1.GetInfoMongoAll)
	v1MongoRouter.GET("/:hostnickname/info", v1.GetInfoMongoHost)
	v1MongoRouter.GET("/:hostnickname/:database/info", v1.GetInfoMongoDatabase)
	v1MongoRouter.GET("/:hostnickname/:database/:collection/info", v1.GetInfoMongoCollection)
	v1MongoRouter.GET("/:hostnickname/:database/:collection/data", v1.GetDataMongoCollection)
	v1MongoRouter.POST("/:hostnickname/:database/:collection/data", v1.StoreDataMongoCollection)

	// Start server
	e.Logger.Fatal(e.Start(":" + s.APISERVER_PORT))
}
