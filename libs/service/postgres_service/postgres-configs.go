package postgres_service

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"reflect"
)

type PostgresConfig struct {
	Nickname           string `json:"nickname"`
	Hostname           string `json:"hostname"`
	Port               int    `json:"port"`
	User               string `json:"user"`
	Password           string `json:"password"`
	Database           string `json:"database"`
	ConnectionPoolSize int    `json:"connectionPoolSize"`
}

func LoadConfig(confFilePath *string) map[string]PostgresConfig {
	var configFilePath string = "./postgres.json"
	if confFilePath != nil {
		configFilePath = *confFilePath
	}

	blob, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		log.Printf(`Failed to read %s\n`, configFilePath)
	}
	bytes := []byte(blob)
	var obj map[string]PostgresConfig
	if err := json.Unmarshal(bytes, &obj); err != nil {
		log.Fatal(err)
	}
	keys := reflect.ValueOf(obj).MapKeys()
	for _, k := range keys {
		key := k.String()
		var conf PostgresConfig = obj[key]
		conf.Nickname = key
		obj[key] = conf
	}
	return obj
}
