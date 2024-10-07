package api

import (
	sync "sync"

	defs "github.com/bbitere/gameapi.git/pkg/defs"
)

const GAME_WILL_START  = 10.0

type ECrashGameState string;
const(

	
	ECrashGameState_Announce  = "announce";
	ECrashGameState_Starting = "starting";
	ECrashGameState_Start    = "start";
	ECrashGameState_Playing  = "playing";
	ECrashGameState_WaitEnd  = "wait";
)
type CrashGameLogic struct{

	mutex					sync.Mutex
	gameState				ECrashGameState		
	gameWillStartInSecond	float64
	targetMultiplicator		defs.TDecimal
}

type CrashUserData struct{

	winCurrHand 		defs.TDecimal;
    moneyBalance 		defs.TDecimal;
    
    //gameData: Array<{multiplCashout:number, bet:number}>;
    bAnnouceFirstTime 	bool;
    timerAnnonce 		float64;
    bUserEndedGame 		bool;
    bUserDidCashout 	bool;
}
func (This *CrashGameLogic) Constr()  *CrashGameLogic{

	This.initGame();
}

func (This*CrashGameLogic) initGame(){

	This.gameState = EGameState_Annonce;
	This.gameWillStartInSecond = GAME_WILL_START
	This.targetMultiplicator = 0;

}
func (This*CrashGameLogic) startPlayingGame(){


}