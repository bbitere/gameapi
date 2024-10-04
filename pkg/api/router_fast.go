package api

import (
	// Swagger middleware
	"github.com/valyala/fasthttp"
	//"encoding/json"
	//"fmt"
	//"net/http"
	//httpSwagger "github.com/swaggo/http-swagger"
)

func RouterFast_Handler(ctx *fasthttp.RequestCtx) {

	switch string(ctx.Path()) {
	case "/authenticate": 	RouterFast_processRoute("Authenticate", Controller_Authenticate )
	case "/play":  			RouterFast_processRoute("Play", Controller_Play ) 
	case "/playautobet":  	RouterFast_processRoute("PlayAutobet", Controller_PlayAutobet ) 
	case "/history": 		RouterFast_processRoute("HistoryList", Controller_HistoryList ) 
	case "/cashout":  		RouterFast_processRoute("Cashout", Controller_Cashout ) 
	case "/gameupdate":  	RouterFast_processRoute("GameUpdate", Controller_GameUpdate ) 
		
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

	RouterFast_ConfigureCORS(ctx)
}

