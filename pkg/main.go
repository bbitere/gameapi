package main

import (
	api "github.com/bbitere/gameapi.git"
	cors "github.com/gin-contrib/cors"
	gin "github.com/gin-gonic/gin"
)


func main() {


	var httprouter  = gin.New()
	var httpsrouter = gin.New()

	var corsConfig = cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowMethods = []string{"GET", "POST"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type"}

	httprouter.Use( cors.New(corsConfig) )
	httpsrouter.Use( cors.New(corsConfig) )

	api.Api_InitRouter(httprouter, httpsrouter);


}