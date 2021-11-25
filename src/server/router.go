package server

import (
	"collect/api/db"
	v1 "collect/api/v1"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	dbConnector := db.GetConnector()

	router := gin.New()
	router.Use(gin.Logger())
	// router.Use(middleware.IsAuthorized())

	collectAPIGroupV1 := router.Group("api/v1")
	{
		collectControllerV1 := v1.NewCollectController(dbConnector)
		collectAPIGroupV1.GET("/getForm", collectControllerV1.GetFormHandler)
		collectAPIGroupV1.GET("/getQuestion", collectControllerV1.GetQuestionHandler)
		collectAPIGroupV1.GET("/getResponse", collectControllerV1.GetResponseHandler)
		collectAPIGroupV1.GET("/getResponseBank", collectControllerV1.GetResponseBankHandler)
	}

	return router
}
