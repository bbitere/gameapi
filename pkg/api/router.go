package api

import (
	gin "github.com/gin-gonic/gin"
	//api "github.com/bbitere/gameapi.git/pkg/api"
	//_ "your_project/docs" // importă pachetul docs generat de Swagger

	"fmt"
	"net/http"

	utils "github.com/bbitere/gameapi.git/pkg/utils"
)

func Api_InitRouter( httpRouter *gin.Engine ){

	httpRouter.POST("/authenticate",  processRoute("Authenticate", Controller_Authenticate ))
	httpRouter.POST("/play",  processRoute("Play", Controller_Play ) )
	httpRouter.POST("/playautobet",  processRoute("PlayAutobet", Controller_PlayAutobet ) )
	httpRouter.POST("/history",  processRoute("HistoryList", Controller_HistoryList ) )
	httpRouter.POST("/cashout",  processRoute("Cashout", Controller_Cashout ) )
	httpRouter.POST("/gameupdate",  processRoute("GameUpdate", Controller_GameUpdate ) )
}

func processRoute[TypeInput any, TypeResponse any, TypeResponseErr any]( 
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