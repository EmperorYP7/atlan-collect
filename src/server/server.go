package server

import (
	"collect/api/config"
	"collect/api/db"
	"collect/api/logger"
)

func Init() {
	logger.Init()
	db.Init()

	r := NewRouter()

	r.Run(":" + config.Get().ServerPort)
}
