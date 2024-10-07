package api

import (
	"sync"

	defs "github.com/bbitere/gameapi.git/pkg/defs"
	models "github.com/bbitere/gameapi.git/pkg/models"
	utils "github.com/bbitere/gameapi.git/pkg/utils"
)

const GAME_WILL_START  = 10.0


type CrashGameLogic struct{

	mutex		sync.Mutex
	GameState 	models.GameState

	PlayersList	[]*models.Player
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

	This.PlayersList = make([]*models.Player, 0, 100);
	This.initGame();
	return This;
}

func (This*CrashGameLogic) initGame(){

	This.GameState.GameState = EGameState_Annonce;
	This.GameState.GameWillStartInSecond = GAME_WILL_START
	This.GameState.TargetMultiplicator = int32(API.Random.GetRandom());
}
func (This*CrashGameLogic) startPlayingGame(){


}

func (This*CrashGameLogic) addPlayer(name string){

	var player = models.Player{
		PlayerName: name,
		IsHidden: false,
		CashoutMultiplicator: 0,
		NextRound_CashoutMultiplicator: 0,
	}
	
	utils.Arr_Append(&This.PlayersList, &player );
}