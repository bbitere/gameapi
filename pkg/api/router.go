package api

import (
	gin "github.com/gin-gonic/gin"
	//api "github.com/bbitere/gameapi.git/pkg/api"
	//_ "your_project/docs" // importă pachetul docs generat de Swagger

	"net/http"

	http "github.com/bbitere/gameapi.git/pkg/http"
)

func Api_InitRouter( httpRouter *gin.Engine ){

	httpRouter.POST("/authenticate",  	http.Http_processRoute("Authenticate", Controller_Authenticate ))
	httpRouter.POST("/play",  			http.Http_processRoute("Play", Controller_Play ) )
	httpRouter.POST("/playautobet",  	http.Http_processRoute("PlayAutobet", Controller_PlayAutobet ) )
	httpRouter.POST("/history",  		http.Http_processRoute("HistoryList", Controller_HistoryList ) )
	httpRouter.POST("/cashout",  		http.Http_processRoute("Cashout", Controller_Cashout ) )
	httpRouter.POST("/gameupdate",  	http.Http_processRoute("GameUpdate", Controller_GameUpdate ) )
}

