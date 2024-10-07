package http

import (
	"fmt"
	"net/http"

	utils "github.com/bbitere/gameapi.git/pkg/utils"
	cors "github.com/gin-contrib/cors"
	gin "github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type CertsHttps struct{
	Cert_pem 	string
	Key_pem		string
}

func Init_HttpServer(
	certsForHttps *CertsHttps, URL_PORT string, 
	fnInitApi func(router* gin.Engine) ){

	var httprouter  = gin.New()
	//var httpsrouter = gin.New()

	var corsConfig = cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowMethods = []string{"GET", "POST"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type"}

	httprouter.Use( cors.New(corsConfig) )
	//httpsrouter.Use( cors.New(corsConfig) )

	fnInitApi(httprouter);
	//api.Api_InitRouter( httprouter);
	//api.Api_InitRouter( httpsrouter);

	//expose the endpoint for swagger documentation
	httprouter.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//install : openssl req -newkey rsa:4096 -nodes -sha256 -keyout key.pem -x509 -days 365 -out cert.pem

	if( certsForHttps != nil ){
		// provide certs: "cert.pem" și "key.pem"
		errRouter := http.ListenAndServeTLS(":"+URL_PORT, 
							certsForHttps.Cert_pem, 
							certsForHttps.Key_pem, httprouter )
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