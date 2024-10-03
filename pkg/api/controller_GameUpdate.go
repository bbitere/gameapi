package api

import (
	defs "github.com/bbitere/gameapi.git/pkg/defs"
)

//"fmt"
//"net/http"

//utils "github.com/bbitere/gameapi.git/pkg/utils"
//gin "github.com/gin-gonic/gin"
//_ "your_project/docs" // importă pachetul docs generat de Swagger

type EGameState string;
const(
	
	EGameState_InitServer="initS"
	EGameState_InitGame="initgame"
	EGameState_Annonce="anounce"
	EGameState_Starting="starting"
	EGameState_Playing="playing"
	EGameState_Msg_WinOrLose ="end"
	EGameState_WaitRestart="wait-restart"	
)

type EBetStatus string;
const(
    EBetStatus_None=""
    EBetStatus_Bet_CurrRound= "Bet1"
    EBetStatus_Bet_NextRound= "Bet2"// 
    EBetStatus_Bet_CurrAndNextRound= "Bet12"// 
)

type ECrashEndLine string;
const(
    ECrashEndLine_End     = "e"// game end, no msg is displaying
    ECrashEndLine_Win     = "w"// win
    ECrashEndLine_Lose    = "l"// plane crash.
)



type TypeCrashArg_Annonce        struct {
	GameWillStart 	defs.TTime	`json:"gameWillStart"`
}

type TypeCrashArg_Playing        struct {
	Timeline defs.TTime				`json:"timeline"`
	Multiplicator 	defs.TDecimal	`json:"multiplicator"`
}
type TypeCrashArg_MsgWinOrLose   struct {
	Timeline 		defs.TTime		`json:"timeline"`
	Multiplicator 	defs.TDecimal	`json:"multiplicator"`
	IsWin			ECrashEndLine	`json:"isWin"`
	WalletAfter 	defs.TDecimal	`json:"walletAfter"`
}



type GameUpdateInput struct{

	SessionID				string 			`json:"sessionUID"`
}
type GameUpdateResponse struct{
	
	GameState 				EGameState		`json:"gameState"`
	BetStatus 				EBetStatus		`json:"betStatus"`
	PlayerList 				[]PlayerItem	`json:"playerlist"`

	StateAnnounce			*TypeCrashArg_Annonce	`json:"announce"`
	StatePlaying			*TypeCrashArg_Playing	`json:"playing"`
	StateMsgWinLose			*TypeCrashArg_MsgWinOrLose	`json:"msg"`

	Error  					int    			`json:"error"`
	Description 			string  		`json:"description"`
}
type GameUpdateResponseErr PlayResponse
// @Summary update the entire game for all user
// @ID Controller_GameUpdate
// @Produce json
// @Param data body GameUpdateInput true "Play input data"
// @Success 200 {object} GameUpdateResponse
// @Failure 400 {object} GameUpdateResponseErr
// @Router /gameupdate [post]
func Controller_GameUpdate(inp *GameUpdateInput) (*GameUpdateResponse, *GameUpdateResponseErr, error){

	var outData = 	GameUpdateResponse{					
					Error: 0,
					Description: "",
				};

	return &outData, nil, nil;
}



