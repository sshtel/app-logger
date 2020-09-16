package utils

import (
	"time"
	"os"
)

func ParseTimeSimple(timeStr string) (time.Time, error) {
	layout := "2006-01-02T15:04:05"
	return time.Parse(layout, timeStr)	
}

func GetOsEnvWithDef(key string, def string) string {
	val := os.Getenv(key)
	if val == "" { return def }
	return val
}
