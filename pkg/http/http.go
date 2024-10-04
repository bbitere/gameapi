package http

import (
	"fmt"
	"net/http"

	api "github.com/bbitere/gameapi.git/pkg/api"
	utils "github.com/bbitere/gameapi.git/pkg/utils"
	cors "github.com/gin-contrib/cors"
	gin "github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Init_HTTP_HTTPS(isHttps bool, URL_PORT string){

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


func Http_processRoute[TypeInput any, TypeResponse any, TypeResponseErr any]( 
	nameMethod string, 
	fnController func(inp* TypeInput) (*TypeResponse, *TypeResponseErr, error),
	
	)  func ( c *gin.Context ) {

	var fnInnerProcess = func( c *gin.Context ){

		var inputData = new (TypeInput);

		utils.Log_log(nameMethod, "", 1, fmt.Sprint(inputData), "" )

		var err = c.BindJSON(&inputData);
		if( err != nil){

			utils.Log_log(nameMethod, "", 1, fmt.Sprint(err), "Cannot read Token" )
			c.IndentedJSON( http.StatusOK, err)
		}
		//-------------------------------------------------------------------------

		var outData, outDataErr, err1 = fnController(inputData);

		//-------------------------------------------------------------------------
		if( err1 != nil){

			if( outDataErr != nil){ }
			c.IndentedJSON(http.StatusOK, outData)
		}else{

			c.IndentedJSON(http.StatusOK, outData)
		}
	}

	
	return fnInnerProcess;
}