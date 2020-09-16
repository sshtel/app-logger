package server

import (
	"net/http"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	model "github.com/sshtel/app-logger/log-server/model"
	utils "github.com/sshtel/app-logger/log-server/utils"
	mongoService "github.com/sshtel/app-logger/log-server/service/mongodb"
	mysqlService "github.com/sshtel/app-logger/log-server/service/mysql"

)
type ApiServer struct {
	APISERVER_PORT string
	mongoRouter *mongoService.MongoRouterService
	mysqlRouter *mysqlService.MysqlRouterService
}

func NewApiServer(mongoRouter *mongoService.MongoRouterService) ApiServer {
	p := new(ApiServer)
	p.APISERVER_PORT = utils.GetOsEnvWithDef("APISERVER_PORT", "8080")
	p.mongoRouter = mongoRouter
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
	e.GET("/api/ping", s.pong)
	e.GET("/api/v1/mongo/info", s.getInfoMongoAll)
	e.GET("/api/v1/mongo/:hostnickname/info", s.getInfoMongoHost)
	e.GET("/api/v1/mongo/:hostnickname/:database/info", s.getInfoMongoDatabase)
	e.GET("/api/v1/mongo/:hostnickname/:database/:collection/info", s.getInfoMongoCollection)
	e.GET("/api/v1/mongo/:hostnickname/:database/:collection/data", s.getDataMongoCollection)
	e.POST("/api/v1/mongo/:hostnickname/:database/:collection/data", s.storeDataMongoCollection)

	// Start server
	e.Logger.Fatal(e.Start(":" + s.APISERVER_PORT))
}


func (s *ApiServer) pong(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}

func (s *ApiServer) getInfoMongoAll(c echo.Context) error {
	result := "resut"
	return c.JSON(http.StatusOK, &result)
}

func (s *ApiServer) getInfoMongoHost(c echo.Context) error {
	hostnickname := c.Param("hostnickname")
	fmt.Println(hostnickname)
	result := "resut"
	return c.JSON(http.StatusOK, &result)
}

func (s *ApiServer) getInfoMongoDatabase(c echo.Context) error {
	hostnickname := c.Param("hostnickname")
	database := c.Param("database")
	fmt.Println(hostnickname)
	fmt.Println(database)

	result := "resut"
	return c.JSON(http.StatusOK, &result)
}

func (s *ApiServer) getInfoMongoCollection(c echo.Context) error {
	hostnickname := c.Param("hostnickname")
	database := c.Param("database")
	collection := c.Param("collection")
	fmt.Println(hostnickname)
	fmt.Println(database)
	fmt.Println(collection)

	result := "resut"
	return c.JSON(http.StatusOK, &result)
}


func (s *ApiServer) getDataMongoCollection(c echo.Context) error {
	hostnickname := c.Param("hostnickname")
	database := c.Param("database")
	collection := c.Param("collection")
	fmt.Println(hostnickname)
	fmt.Println(database)
	fmt.Println(collection)

	result := "resut"
	return c.JSON(http.StatusOK, &result)
}


func (s *ApiServer) storeDataMongoCollection(c echo.Context) error {
	hostnickname := c.Param("hostnickname")
	database := c.Param("database")
	collection := c.Param("collection")
	from := c.QueryParam("from")
	to := c.QueryParam("to")
	fmt.Println(hostnickname)
	fmt.Println(database)
	fmt.Println(collection)
	fmt.Println(from)
	fmt.Println(to)

	result := "resut"

	json_map := echo.Map{}
	if err := c.Bind(&json_map); err != nil { return err }
	fmt.Println(json_map)
	s.mongoRouter.GetInputChannel("default") <- model.LogData{
		Timestamp: "time",
		JsonBody: json_map,
		QueryParam: json_map,
	}

	return c.JSON(http.StatusOK, &result)
}


func (s *ApiServer) storeToMongoDbCollection(c echo.Context) error {
	json_map := echo.Map{}
	if err := c.Bind(&json_map); err != nil {
		return err
	}

	timestamp := json_map["timestamp"].(string)
	data := json_map["data"].(string)
	fmt.Println(timestamp)
	fmt.Println(data)

	result := "result"

	return c.JSON(http.StatusOK, &result)
}
