package api

import (
	gin "github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	//_ "your_project/docs" // importă pachetul docs generat de Swagger
)

// @title InteractiveService
// @version 1.0
// @description This is the backend for the Interactive service
// @host localhost:8081
// @BasePath /
// @query.collection.format multi
func Api_InitRouter( httpRouter *gin.Engine, httpsRouter*gin.Engine ){

	//router := gin.Default()

	//router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	httpRouter.POST("/authenticate", Authenticate)
	httpsRouter.POST("/authenticate", Authenticate)


	httpRouter.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	

}