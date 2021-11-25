package db

import (
	"database/sql"
	"sync"
	"collect/api/logger"
	"collect/api/config"
	_ "github.com/go-sql-driver/mysql"
)

var connector *sql.DB
var once sync.Once

type DBConfig struct {
	DBUserName           string
	DBPassword           string
	DBHost               string
	DBPort               string
	DBName               string
}

func NewDBClient(config *config.Config) *sql.DB {
	url := config.DBUserName + ":" + config.DBPassword + "@tcp(" + config.DBHost + ":" + config.DBPort + ")/" + config.DBName
	client, err := sql.Open("mysql", url)

	if err != nil {
		logger.Client().Fatal("Couldn't connect to DB!")
		logger.Client().Error(err.Error())
		panic(err.Error())
	}

	return client
}

func Init() {
	once.Do(func() {
		config := config.Get()
		connector = NewDBClient(&config)
		logger.Client().Info("DB is connected")
	})
}

func GetConnector() *sql.DB {
	return connector
}
