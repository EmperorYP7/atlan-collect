package server

import (
	"fmt"
	"collect/api/db"
	"collect/api/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	dbConnector := db.GetConnector()

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(middleware.IsAuthorized())

	collectAPIGroupV1 := router.Group("api/v1") {
		
	}
}