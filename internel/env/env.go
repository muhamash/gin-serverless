package env

import (
	"log"
	"os"
	"strconv"
)

func GetEnvString(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("Missing required env var: %s", key)
	}
	return val
}

func GetEnvInt(key string) int {
	val := GetEnvString(key)
	num, err := strconv.Atoi(val)
	if err != nil {
		log.Fatalf("Invalid int for env var %s: %v", key, err)
	}
	return num
}