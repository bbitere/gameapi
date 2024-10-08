package api

import (
	// Swagger middleware
	http1 "github.com/bbitere/gameapi.git/pkg/http"
	"github.com/valyala/fasthttp"
	//"encoding/json"
	//"fmt"
	//"net/http"
	//httpSwagger "github.com/swaggo/http-swagger"
)

func RouterFast_Handler(ctx *fasthttp.RequestCtx) {

	switch string(ctx.Path()) {
	case "/authenticate": 	http1.HttpFast_processRoute("Authenticate", Controller_Authenticate )
	case "/play":  			http1.HttpFast_processRoute("Play", Controller_Play ) 
	case "/playautobet":  	http1.HttpFast_processRoute("PlayAutobet", Controller_PlayAutobet ) 
	case "/history": 		http1.HttpFast_processRoute("HistoryList", Controller_HistoryList ) 
	case "/cashout":  		http1.HttpFast_processRoute("Cashout", Controller_Cashout ) 
	case "/gameupdate":  	http1.HttpFast_processRoute("GameUpdate", Controller_GameUpdate ) 
		
	case "/swagger":

		/*
		//httpSwagger.Handler(ctx.Response.BodyWriter(), ctx.Request) // Răspunsul se face direct
		// Apelează handler-ul Swagger folosind `http.ResponseWriter` și `http.Request`
		r := &http.Request{
			Method: string(ctx.Method()),
			Header: ctx.Request.Header,
			Body:   http.NoBody,
			// Poți adăuga mai multe câmpuri după cum este necesar
		}
	
		w := &http.Response{
			Header: http.Header{},
			// Poți adăuga mai multe câmpuri după cum este necesar
		}
	
		// Redirecționează cererea către handler-ul Swagger
		httpSwagger.Handler.ServeHTTP(w, r)
		// Scrie răspunsul în contextul fasthttp
		ctx.SetStatusCode(w.StatusCode)
		w.Header.Write(ctx.Response.BodyWriter())
	
		ctx.Redirect(  "/swagger/", fasthttp.StatusSeeOther );
			// &ctx.Response, &ctx.Request, "/swagger/", fasthttp.StatusSeeOther )
		*/
	default:
		ctx.Error("Not Found", fasthttp.StatusNotFound)
	}

	http1.HttpFast_ConfigureCORS(ctx)
}

