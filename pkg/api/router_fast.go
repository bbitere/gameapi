package api

import (
	// Swagger middleware
	"github.com/valyala/fasthttp"

	"encoding/json"
	"fmt"

	//"net/http"

	//httpSwagger "github.com/swaggo/http-swagger"

	utils "github.com/bbitere/gameapi.git/pkg/utils"
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

func RouterFast_processRoute[TypeInput any, TypeResponse any, TypeResponseErr any]( 
	nameMethod string, 
	fnController func(inp* TypeInput) (*TypeResponse, *TypeResponseErr, error),

	)  func ( ctx *fasthttp.RequestCtx ) {

	var fnInnerProcess = func( ctx *fasthttp.RequestCtx ){

		if string(ctx.Method()) == "POST" {

			var inputData = new (TypeInput);

			utils.Log_log(nameMethod, "", 1, fmt.Sprint(inputData), "" )

			// Deserializare JSON
			if err := json.Unmarshal(ctx.PostBody(), &inputData); err != nil {

				utils.Log_log(nameMethod, "", 1, fmt.Sprint(err), "Cannot read Token" )
				ctx.Error(err.Error(), fasthttp.StatusBadRequest)
				return
			}
			
			//-------------------------------------------------------------------------

			var outData, outDataErr, err1 = fnController(inputData);

			//-------------------------------------------------------------------------
			if( err1 != nil){

				if( outDataErr != nil){ }
				response, _ := json.Marshal(outDataErr)

				ctx.SetContentType("application/json")
				ctx.SetStatusCode(fasthttp.StatusOK)
				ctx.Write(response)
			}else{

				response, _ := json.Marshal(outData)

				ctx.SetContentType("application/json")
				ctx.SetStatusCode(fasthttp.StatusOK)
				ctx.Write(response)
			}
		}else {
			ctx.Error("Method Not Allowed", fasthttp.StatusMethodNotAllowed)
		}
	}
	
	return fnInnerProcess;
}

func RouterFast_ConfigureCORS(ctx *fasthttp.RequestCtx){
	
	ctx.Response.Header.Set("Access-Control-Allow-Origin", "*") // Permite toate originile
    ctx.Response.Header.Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS") // Permite metodele
    ctx.Response.Header.Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Authorization") // Permite antetele

}

/*
// NewFastHTTPHandler wraps net/http handler to fasthttp request handler,
// so it can be passed to fasthttp server.
//
// While this function may be used for easy switching from net/http to fasthttp,
// it has the following drawbacks comparing to using manually written fasthttp
// request handler:
//
//   - A lot of useful functionality provided by fasthttp is missing
//     from net/http handler.
//   - net/http -> fasthttp handler conversion has some overhead,
//     so the returned handler will be always slower than manually written
//     fasthttp handler.
//
// So it is advisable using this function only for quick net/http -> fasthttp
// switching. Then manually convert net/http handlers to fasthttp handlers
// according to https://github.com/valyala/fasthttp#switching-from-nethttp-to-fasthttp .
func NewFastHTTPHandler(h http.Handler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		var r http.Request
		if err := ConvertRequest(ctx, &r, true); err != nil {
			ctx.Logger().Printf("cannot parse requestURI %q: %v", r.RequestURI, err)
			ctx.Error("Internal Server Error", fasthttp.StatusInternalServerError)
			return
		}

		w := netHTTPResponseWriter{w: ctx.Response.BodyWriter()}
		h.ServeHTTP(&w, r.WithContext(ctx))

		ctx.SetStatusCode(w.StatusCode())
		haveContentType := false
		for k, vv := range w.Header() {
			if k == fasthttp.HeaderContentType {
				haveContentType = true
			}

			for _, v := range vv {
				ctx.Response.Header.Add(k, v)
			}
		}
		if !haveContentType {
			// From net/http.ResponseWriter.Write:
			// If the Header does not contain a Content-Type line, Write adds a Content-Type set
			// to the result of passing the initial 512 bytes of written data to DetectContentType.
			l := 512
			b := ctx.Response.Body()
			if len(b) < 512 {
				l = len(b)
			}
			ctx.Response.Header.Set(fasthttp.HeaderContentType, http.DetectContentType(b[:l]))
		}
	}
}

*/