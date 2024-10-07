package api

import (
	//"fmt"
	//"net/http"

	defs "github.com/bbitere/gameapi.git/pkg/defs"
	//utils "github.com/bbitere/gameapi.git/pkg/utils"
	//gin "github.com/gin-gonic/gin"
	//_ "your_project/docs" // importă pachetul docs generat de Swagger
)


type PlayAutobetInput struct{

	SessionID				string 			`json:"sessionUID"`
	PlayerToken				string 			`json:"playerToken"`
	
	Bet 					defs.TDecimal	`json:"bet"`
	MultiplicatorCashout 	defs.TDecimal	`json:"multiplBashout"`
    NumOfAutoBets 			int				`json:"numOfAutobet"`
	StopOnProfit 			defs.TDecimal	`json:"stopOnProfit"`
    StopOnLoss 				defs.TDecimal	`json:"StopOnLoss"`
}
type PlayAutobetResponse struct{

	Error  					int    			`json:"error"`
	Description 			string  		`json:"description"`
}
type PlayAutobetResponseErr PlayResponse
// @Summary Play user
// @ID Controller_PlayAutobet
// @Produce json
// @Param data body PlayAutobetInput true "Play input data"
// @Success 200 {object} PlayAutobetResponse
// @Failure 400 {object} PlayAutobetResponseErr
// @Router /playautobet [post]
func Controller_PlayAutobet(inp *PlayAutobetInput) (*PlayAutobetResponse, *PlayAutobetResponseErr, error){

	var outData = 	PlayAutobetResponse{					
					Error: 0,
					Description: "",
				};

	return &outData, nil, nil;
}



