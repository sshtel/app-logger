package defs

import (
	"os"
)
func GetenvWithDef(key string, def string) string {
	val := os.Getenv(key)
	if val == "" { return def }
	return val
}


var (
	CONFIG_MONGO = ""
)

func LoadEnvs() {
	CONFIG_MONGO = GetenvWithDef("CONFIG_MONGO", "./configs/mongo.json")
}