package api

import (
	//"fmt"
	//"net/http"

	defs "github.com/bbitere/gameapi.git/pkg/defs"
	//utils "github.com/bbitere/gameapi.git/pkg/utils"
	//gin "github.com/gin-gonic/gin"
	//_ "your_project/docs" // importă pachetul docs generat de Swagger
)

type ECommand = string;
const (

	Command_Default = "def";
	Command_BetNextRound = "next";
	Command_Cancel = "cancel";
)

type PlayInput struct{

	SessionID				string 			`json:"sessionUID"`
	Bet 					defs.TDecimal	`json:"bet"`
	MultiplicatorCashout 	defs.TDecimal	`json:"multiplBashout"`
	Command					ECommand		`json:"commandID"`
}
type PlayResponse struct{

	Wallet					defs.TDecimal	`json:"wallet"`
	Error  					int    			`json:"error"`
	Description 			string  		`json:"description"`
}
type PlayResponseErr PlayResponse
// @Summary Play user
// @ID Controller_Play
// @Produce json
// @Param data body PlayInput true "Play input data"
// @Success 200 {object} PlayResponse
// @Failure 400 {object} PlayResponseErr
// @Router /play [post]
func Controller_Play(inp *PlayInput) (*PlayResponse, *PlayResponseErr, error){

	var outData = 	PlayResponse{					
					Error: 0,
					Description: "",
				};

	return &outData, nil, nil;
}



