package main

import (

	//api "github.com/bbitere/gameapi.git/pkg/api"
	"github.com/bbitere/gameapi.git/pkg/api"
	http "github.com/bbitere/gameapi.git/pkg/http"
	utils "github.com/bbitere/gameapi.git/pkg/utils"
	"github.com/gin-gonic/gin"

	//_ "github.com/bbitere/gameapi.git/docs"
	_ "github.com/bbitere/gameapi.git/pkg/api"
	_ "github.com/bbitere/gameapi.git/pkg/docs"
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

	defer utils.LogPanic();
	utils.StartLogPanic();

	utils.Log_initContainerLogger( );

	var server = api.DataServer_init();

	http.Init_HttpServer(nil, "8080", 
		func(router *gin.Engine){
			api.Api_InitRouter(router);
		} );

	
		
	/*
	var certsForHttps = http.CertsHttps{
		Cert_pem :  "cert.pem",
		Key_pem : "key.pem",
	}
	http.Init_HttpServer(&certsForHttps, "8080", 
		func(router *gin.Engine){
			api.Api_InitRouter(router);
		} );
	//*/

	/*
	http.Init_HTTPFast(nil, "8080",
		func() func(ctx *fasthttp.RequestCtx){
			return api.RouterFast_Handler;
		} );

	var certsForFastHttps = http.CertsFastHttps{
			Server_crt :  "server.crt",
			Server_key : "server.key",
		}

	http.Init_HTTPFast(&certsForFastHttps, "8080",
		func() func(ctx *fasthttp.RequestCtx){
			return api.RouterFast_Handler;
		} );
	//*/


	server.GameLogic.InitMainThread( server );
	
}



