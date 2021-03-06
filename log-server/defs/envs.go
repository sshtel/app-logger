package defs

import (
	"os"
)

func GetenvWithDef(key string, def string) string {
	val := os.Getenv(key)
	if val == "" {
		return def
	}
	return val
}

var (
	CONFIG_POSTGRES = ""
	CONFIG_MONGO    = ""
)

func LoadEnvs() {
	CONFIG_POSTGRES = GetenvWithDef("CONFIG_DB", "./configs/postgres.json")
	CONFIG_MONGO = GetenvWithDef("CONFIG_MONGO", "./configs/mongo.json")
}
