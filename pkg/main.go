package main

import (
	"fmt"
	"net/http"

	api "github.com/bbitere/gameapi.git/pkg/api"
	utils "github.com/bbitere/gameapi.git/pkg/utils"
	cors "github.com/gin-contrib/cors"
	gin "github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	//_ "github.com/bbitere/gameapi.git/docs"
	_ "github.com/bbitere/gameapi.git/pkg/api"
	_ "github.com/bbitere/gameapi.git/pkg/docs"

	"github.com/valyala/fasthttp"
)

//url of swagger: http://localhost:8081/docs/swagger/index.html
const URL_PORT = "8081"

// @title InteractiveService
// @version 1.0
// @description This is the backend for the Interactive service
// @host localhost:8081
// @BasePath /
// @query.collection.format multi
func main() {

	utils.Log_initContainerLogger( );

	init_HTTP_HTTPS(false, "8080" );
	//init_HTTP_HTTPS(true, "8443" );

	init_HTTPFast(false, "8080" );
}

func init_HTTP_HTTPS(isHttps bool, URL_PORT string){

	var httprouter  = gin.New()
	//var httpsrouter = gin.New()

	var corsConfig = cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowMethods = []string{"GET", "POST"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type"}

	httprouter.Use( cors.New(corsConfig) )
	//httpsrouter.Use( cors.New(corsConfig) )

	api.Api_InitRouter( httprouter);
	//api.Api_InitRouter( httpsrouter);

	//expose the endpoint for swagger documentation
	httprouter.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//install : openssl req -newkey rsa:4096 -nodes -sha256 -keyout key.pem -x509 -days 365 -out cert.pem

	if( isHttps ){
		// provide certs: "cert.pem" și "key.pem"
		errRouter := http.ListenAndServeTLS(":"+URL_PORT, "cert.pem", "key.pem", httprouter )
		if errRouter != nil {
			utils.Log_log("routner not init", "", 3, fmt.Sprint(errRouter), "")
		}
	}else{
		var errRouter = http.ListenAndServe(":"+ URL_PORT, httprouter)
		if( errRouter != nil){
			utils.Log_log("routner not init", "", 3, fmt.Sprint(errRouter), "")
		}
	}
}

func init_HTTPFast(bHttps bool, URL_PORT string){

	if( bHttps){

		fmt.Println("Server running on https://localhost:"+PORT)
		if err := fasthttp.ListenAndServeTLS(":"+PORT, "server.crt", "server.key", api.RouterFast_Handler); err != nil {
			panic(err)
		}

	} else {
		// Rulează serverul pe portul 8080
		fmt.Println("Server running on http://localhost:" + PORT)
		if err := fasthttp.ListenAndServe(":"+PORT, api.RouterFast_Handler); err != nil {
			panic(err)
		}
	}

}