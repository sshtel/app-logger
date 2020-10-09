package handler_v1
import (
	"fmt"
	"net/http"
	"github.com/labstack/echo"
	mongoService "../../service/mongodb"
	model "../../model"
)

func GetInfoMongoAll(c echo.Context) error {
	result := "resut"
	return c.JSON(http.StatusOK, &result)
}

func GetInfoMongoHost(c echo.Context) error {
	hostnickname := c.Param("hostnickname")
	fmt.Println(hostnickname)
	result := "resut"
	return c.JSON(http.StatusOK, &result)
}

func GetInfoMongoDatabase(c echo.Context) error {
	hostnickname := c.Param("hostnickname")
	database := c.Param("database")
	fmt.Println(hostnickname)
	fmt.Println(database)

	result := "resut"
	return c.JSON(http.StatusOK, &result)
}

func GetInfoMongoCollection(c echo.Context) error {
	hostnickname := c.Param("hostnickname")
	database := c.Param("database")
	collection := c.Param("collection")
	fmt.Println(hostnickname)
	fmt.Println(database)
	fmt.Println(collection)

	result := "resut"
	return c.JSON(http.StatusOK, &result)
}


func GetDataMongoCollection(c echo.Context) error {
	hostnickname := c.Param("hostnickname")
	database := c.Param("database")
	collection := c.Param("collection")
	fmt.Println(hostnickname)
	fmt.Println(database)
	fmt.Println(collection)

	result := "resut"
	return c.JSON(http.StatusOK, &result)
}


func StoreDataMongoCollection(c echo.Context) error {
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
	mongoService.MongoServiceRef.GetInputChannel("default") <- model.LogData{
		Timestamp: "time",
		JsonBody: json_map,
		QueryParam: json_map,
	}

	return c.JSON(http.StatusOK, &result)
}


func StoreToMongoDbCollection(c echo.Context) error {
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
