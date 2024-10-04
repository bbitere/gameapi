package main

import (

	//api "github.com/bbitere/gameapi.git/pkg/api"
	http "github.com/bbitere/gameapi.git/pkg/http"
	utils "github.com/bbitere/gameapi.git/pkg/utils"

	//_ "github.com/bbitere/gameapi.git/docs"
	_ "github.com/bbitere/gameapi.git/pkg/api"
	_ "github.com/bbitere/gameapi.git/pkg/docs"
	//"github.com/valyala/fasthttp"
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

	http.Init_HTTP_HTTPS(false, "8080" );
	//init_HTTP_HTTPS(true, "8443" );

	http.Init_HTTPFast(false, "8080" );
}



