package api

import (
	"fmt"

	defs "github.com/bbitere/gameapi.git/pkg/defs"
	"github.com/bbitere/gameapi.git/pkg/utils"
)

//"fmt"
//"net/http"

//utils "github.com/bbitere/gameapi.git/pkg/utils"
//gin "github.com/gin-gonic/gin"
//_ "your_project/docs" // importă pachetul docs generat de Swagger


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
	Timeline        defs.TTime				`json:"timeline"`
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
	PlayerToken				string 			`json:"playerToken"`
}
type GameUpdateResponse struct{
	
	GameState 				ECrashGameState	`json:"gameState"`
	BetStatus 				EBetStatus		`json:"betStatus"`
	PlayerList 				[]*PlayerItem	`json:"playerlist"`

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

	var This = API;

	This.LogicMutex.Lock()
	defer This.LogicMutex.Unlock()

	var outData = 	GameUpdateResponse{					
		Error: 0,
		Description: "",
	};

	var gameState = &This.GameLogic.GameState;

	var myPlayerInst = utils.SyncArr_Where( &This.GameLogic.PlayersList, nil,
						func(x *PlayerData) bool { return x.model.Token == inp.PlayerToken} );
	if( myPlayerInst == nil){

		var outDataErr = GameUpdateResponse{					

			Error: defs.Error_playerTokenNotFound,
			Description: fmt.Sprintf("Player not found"),
		};
		return &outDataErr, nil, nil;
	}
	
	outData.GameState = ECrashGameState( gameState.GameState )	
	outData.BetStatus = getBetStatus( myPlayerInst );
	outData.PlayerList= utils.SyncArr_Select( &This.GameLogic.PlayersList, 
						func(x *PlayerData) *PlayerItem{
							return &PlayerItem{
							
								Time:    		x.model.TimeStamp,
								PlayerToken:    x.model.Token,								
								PlayerName:   	x.model.PlayerName,
								Cashout:    	defs.Number2Dec( 0 ),
								Bet:     		defs.Number2Dec( 0 ),
								Currency:		"RON",
								Flg:        	EInteractivUserFlag_List,
								IsCashedout: 	false,
							}
						});

	if( !defs.AppConfig_USE_Websocket &&
		myPlayerInst.winOrLose != nil ){

		outData.StateMsgWinLose = &TypeCrashArg_MsgWinOrLose{ 			
			Timeline :     myPlayerInst.winOrLose.Timeline,
			Multiplicator: myPlayerInst.winOrLose.Multiplicator,
			IsWin: myPlayerInst.winOrLose.IsWin,
			WalletAfter: myPlayerInst.winOrLose.WalletAfter,
		}
		//remove the message
		myPlayerInst.winOrLose = nil;
	}else{

		if( gameState.GameState == ECrashGameState_Announce){

			outData.StateAnnounce = &TypeCrashArg_Annonce{ 
				GameWillStart: defs.TTime( gameState.TargetTime - gameState.StateTimer),
			}
		}else
		if( gameState.GameState == ECrashGameState_Starting){

			
		}else
		if( gameState.GameState == ECrashGameState_Playing){

			outData.StatePlaying = &TypeCrashArg_Playing{ 			
				Timeline :     defs.TTime( gameState.TargetTime ),
				Multiplicator: defs.Number2Dec( defs.TNumber( gameState.TargetMultiplicator) ),
			}
			
		}else
		if( gameState.GameState == ECrashGameState_WaitEnd){

			
		}
	}

	return &outData, nil, nil;
}

func getBetStatus( player *PlayerData) EBetStatus {

	if( player.model.CashoutMultiplicator > 0 ){

		if( player.model.NextRound_CashoutMultiplicator > 0 ) {
			return EBetStatus_Bet_CurrAndNextRound
		}else{
			return EBetStatus_Bet_CurrRound
		}
	}else {

		if( player.model.NextRound_CashoutMultiplicator > 0 ) {
			return EBetStatus_Bet_NextRound
		}else {
			return EBetStatus_None
		}
	}
}



