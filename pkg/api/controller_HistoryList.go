package api

import (
	//"fmt"
	//"net/http"

	"time"

	defs "github.com/bbitere/gameapi.git/pkg/defs"
	//utils "github.com/bbitere/gameapi.git/pkg/utils"
	//gin "github.com/gin-gonic/gin"
	//_ "your_project/docs" // importă pachetul docs generat de Swagger
)



type HistoryItem struct{

	Time    		time.Time	`json:"time"`
	Multplicator 	defs.TDecimal	`json:"multiplicator"`  
	Bet     		defs.TDecimal	`json:"bet"`
	BetCurrency		defs.TDecimal	`json:"betCurrency"`
	Details 		string			`json:"details"`
	Currency		string			`json:"currency"`
	Win     		defs.TDecimal	`json:"win"`
}


type HistoryListInput struct{

	SessionID				string 			`json:"sessionUID"`
    StartOffset 			int				`json:"startOffset"`
	Size 					int				`json:"size"`
	
}
type HistoryListResponse struct{

	List                    []HistoryItem	`json:"list"`
	Error  					int    			`json:"error"`
	Description 			string  		`json:"description"`
}
type HistoryListResponseErr PlayResponse
// @Summary Play user
// @ID Controller_HistoryList
// @Produce json
// @Param data body HistoryListInput true "Play input data"
// @Success 200 {object} HistoryListResponse
// @Failure 400 {object} HistoryListResponseErr
// @Router /history [post]
func Controller_HistoryList(inp *HistoryListInput) (*HistoryListResponse, *HistoryListResponseErr, error){

	var outData = 	HistoryListResponse{					
					Error: 0,
					Description: "",
				};

	return &outData, nil, nil;
}



