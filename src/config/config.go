package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv               string
	ServerPort           string
	DBUserName           string
	DBPassword           string
	DBHost               string
	DBPort               string
	DBName               string
}

var config Config

func init() {
	appEnv := os.Getenv("APP_ENV")
	if len(appEnv) == 0 {
		appEnv = "dev"
	}

	configFilePath := ".env"

	switch appEnv {
	case "production":
		configFilePath = ".env.prod"
		break
	case "stage":
		configFilePath = ".env.stage"
		break
	}

	err := godotenv.Load(configFilePath)
	if err != nil {
		fmt.Errorf("", err.Error())
		panic(err.Error())
	}
	config.AppEnv = appEnv
	config.ServerPort = os.Getenv("SERVER_PORT")
	config.DBUserName = os.Getenv("DB_USERNAME")
	config.DBPassword = os.Getenv("DB_PASSWORD")
	config.DBHost = os.Getenv("DB_HOST")
	config.DBPort = os.Getenv("DB_PORT")
	config.DBName = os.Getenv("DB_NAME")
}

func Get() Config {
	return config
}

func IsProduction() bool {
	return config.AppEnv == "production"
}
