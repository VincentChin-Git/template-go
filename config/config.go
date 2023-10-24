package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL  string
	DatabaseName string
	RedisAddr    string
	RedisPass    string
	TokenSecret  string
}

func GetConfig() Config {
	err := godotenv.Load(".env.local")

	if err != nil {
		panic(err)
	}

	databaseURL := os.Getenv("DatabaseURL")
	databaseName := os.Getenv("DatabaseName")
	redisAddr := os.Getenv("RedisAddr")
	redisPass := os.Getenv("RedisPass")
	tokenSecret := os.Getenv("TokenSecret")

	temp := Config{DatabaseURL: databaseURL, DatabaseName: databaseName, RedisAddr: redisAddr, RedisPass: redisPass, TokenSecret: tokenSecret}

	return temp
}
