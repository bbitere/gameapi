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

type EInteractivUserFlag string
const
(
    List = "l"
    Add = "a"
    Del = "d"
)


type PlayerItem struct{

	Time    		time.Time		`json:"time"`
	UserName   		string			`json:"username"`
	Cashout    		defs.TDecimal	`json:"cashout"`
	Bet     		defs.TDecimal	`json:"bet"`
	Currency		string			`json:"currency"`
	Flg        		EInteractivUserFlag `json:"flg"`
	IsCashedout 	bool			`json:"isCashedout"`
}


type PlayerListInput struct{

	SessionID				string 			`json:"sessionUID"`
    StartOffset 			int				`json:"startOffset"`
	Size 					int				`json:"size"`
	
}
type PlayerListResponse struct{

	List                    []PlayerItem	`json:"list"`
	Error  					int    			`json:"error"`
	Description 			string  		`json:"description"`
}
type PlayerListResponseErr PlayResponse
// @Summary Play user
// @ID Controller_PlayerList
// @Produce json
// @Param data body PlayerListInput true "Play input data"
// @Success 200 {object} PlayerListResponse
// @Failure 400 {object} PlayerListResponseErr
// @Router /Player [post]
func Controller_PlayerList(inp *PlayerListInput) (*PlayerListResponse, *PlayerListResponseErr, error){

	var outData = 	PlayerListResponse{					
					Error: 0,
					Description: "",
				};

	return &outData, nil, nil;
}



