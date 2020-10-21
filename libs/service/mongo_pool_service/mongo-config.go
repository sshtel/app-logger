package mongo_pool_service

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"reflect"
)

type MongoConfig struct {
	Nickname           string `json:"nickname"`
	Hostname           string `json:"hostname"`
	Port               int    `json:"port"`
	User               string `json:"user"`
	Password           string `json:"password"`
	Database           string `json:"database"`
	ConnectionPoolSize int    `json:"connectionPoolSize"`
}

func LoadConfig(filePath *string) map[string]MongoConfig {

	var confFilePath string = "./mongo.json"
	if filePath != nil {
		confFilePath = *filePath
	}

	blob, err := ioutil.ReadFile(confFilePath)
	if err != nil {
		log.Printf(`Failed to read %s\n`, confFilePath)
	}
	bytes := []byte(blob)
	var obj map[string]MongoConfig
	if err := json.Unmarshal(bytes, &obj); err != nil {
		log.Fatal(err)
	}
	keys := reflect.ValueOf(obj).MapKeys()
	for _, k := range keys {
		key := k.String()
		var conf MongoConfig = obj[key]
		conf.Nickname = key
		obj[key] = conf
	}
	return obj
}
