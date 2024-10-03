package api

import (
	//"fmt"
	//"net/http"

	defs "github.com/bbitere/gameapi.git/pkg/defs"
	//utils "github.com/bbitere/gameapi.git/pkg/utils"
	//gin "github.com/gin-gonic/gin"
	//_ "your_project/docs" // importă pachetul docs generat de Swagger
)


type CashoutInput struct{

	SessionID				string 			`json:"sessionUID"`
	Bet 					defs.TDecimal	`json:"bet"`
	MultiplicatorCashout 	defs.TDecimal	`json:"multiplBashout"`
    NumOfAutoBets 			int				`json:"numOfAutobet"`
	StopOnProfit 			defs.TDecimal	`json:"stopOnProfit"`
    StopOnLoss 				defs.TDecimal	`json:"StopOnLoss"`
}
type CashoutResponse struct{

	Error  					int    			`json:"error"`
	Description 			string  		`json:"description"`
}
type CashoutResponseErr PlayResponse
// @Summary do cashot for user
// @ID Controller_Cashout
// @Produce json
// @Param data body CashoutInput true "Play input data"
// @Success 200 {object} CashoutResponse
// @Failure 400 {object} CashoutResponseErr
// @Router /cashout [post]
func Controller_Cashout(inp *CashoutInput) (*CashoutResponse, *CashoutResponseErr, error){

	var outData = 	CashoutResponse{					
					Error: 0,
					Description: "",
				};

	return &outData, nil, nil;
}



