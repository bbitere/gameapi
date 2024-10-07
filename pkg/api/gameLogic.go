package api

import (
	"fmt"
	"math"
	"time"

	defs "github.com/bbitere/gameapi.git/pkg/defs"
	models "github.com/bbitere/gameapi.git/pkg/models"
	utils "github.com/bbitere/gameapi.git/pkg/utils"
	uuid "github.com/google/uuid"
)

const DURATION_ANNOUNCE  = 10.0
const DURATION_STARTING  = 3.0
const DURATION_WAIT      = 3.0


const MPLX_20_TIME          = 8000;// I want to have 2.0x at 12 sec
const MPLX_MAX_TIME_OFFSET  = 80000; // 80 secs -> mpl =  60.x 
const MPLX_MAX_MPLX         = 60;    // mpl =  60.x 

type ECrashGameState string

const (
	ECrashGameState_Announce = "announce"
	ECrashGameState_Starting = "starting"
	ECrashGameState_Start    = "start"
	ECrashGameState_Playing  = "playing"
	ECrashGameState_WaitEnd  = "wait"
)


type CrashGameLogic struct{

	//mutex			sync.Mutex
	LastTimeUpdate  time.Time
	GameState 		models.GameState

	PlayersList		utils.SyncArr[*models.Player]
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

	This.PlayersList.Arr = make([]*models.Player, 0, 100);
	//This.initGame();
	return This;
}

func (This*CrashGameLogic) initGame(server *DataServer ){

	This.setState( server, ECrashGameState_Announce );
	
}
func (This*CrashGameLogic) startPlayingGame(){


}

func (This*CrashGameLogic) addPlayer(name string, bIsHidden bool, currency string, bForceSameName bool ) (string, error){

	var foundPlayer = utils.SyncArr_Where( &This.PlayersList, nil,
			func(x *models.Player) bool { return x.PlayerName == name});

	if( foundPlayer != nil && !bForceSameName ){

		return "", fmt.Errorf("Player already found");

	}else{

		var player = models.Player{
			PlayerName: name,
			IsHidden: false,
			Currency: currency,
			CashedoutAtValue: 0,
			CashoutMultiplicator: 0,
			NextRound_CashoutMultiplicator: 0,
			Token: uuid.New(),
		}
		
		utils.SyncArr_Append( &This.PlayersList, &player );

		return player.Token, nil;
	}
}




func (This*CrashGameLogic) notify_AllPlayersGameEnd(serverData* DataServer) {


}

func (This*CrashGameLogic) computeMultiplayerByTime(stepTimerMs float64) float64{

	var currentTime     = stepTimerMs;
	var t               = (currentTime/MPLX_20_TIME);
	var MPLTX_POW       = math.Log(MPLX_MAX_MPLX)/math.Log(MPLX_MAX_TIME_OFFSET/MPLX_20_TIME);
	var multiplicator   = 1.0 + math.Pow( t, MPLTX_POW) + 0.15 * math.Sin( t * math.Pi);

	return multiplicator
}

func (This*CrashGameLogic) setState( serverData* DataServer, newState ECrashGameState){

	var gameState = &This.GameState;
	gameState.StateTimer = 0;

	if( gameState.GameState == ECrashGameState_Announce){

		gameState.TargetTime = DURATION_ANNOUNCE;		
	}else
	if( gameState.GameState == ECrashGameState_Starting){

		gameState.TargetTime = DURATION_STARTING;
	}else
	if( gameState.GameState == ECrashGameState_Playing){

		gameState.TargetTime 			= serverData.Random.GetRandom2(0, 100)
		var mplx =  This.computeMultiplayerByTime( gameState.TargetTime );
		gameState.TargetMultiplicator 	= defs.Float2Number( mplx );
		
	}else
	if( gameState.GameState == ECrashGameState_WaitEnd){

		gameState.TargetTime = DURATION_WAIT;
	}

}

func (This*CrashGameLogic) updateLoop(serverData* DataServer){

	serverData.LogicMutex.Lock()
	defer serverData.LogicMutex.Unlock()

	var gameState = &This.GameState;
	var nowTimer = time.Now()
	var diff = nowTimer.Sub( This.LastTimeUpdate )
	This.LastTimeUpdate = nowTimer;

	gameState.StateTimer += diff.Seconds();

	if( gameState.GameState == ECrashGameState_Announce){

		if( gameState.StateTimer > gameState.TargetTime){
			
			This.setState( serverData, ECrashGameState_Starting );
		}
	}else
	if( gameState.GameState == ECrashGameState_Starting){

		if( gameState.StateTimer > gameState.TargetTime){
			
			This.setState( serverData, ECrashGameState_Playing );
		}

	}else
	if( gameState.GameState == ECrashGameState_Playing){

		if( gameState.StateTimer > gameState.TargetTime){
			
			This.setState( serverData, ECrashGameState_WaitEnd );

			//notify all players that they lose.
			This.notify_AllPlayersGameEnd( serverData );
			
		}else{

			var mplx =  This.computeMultiplayerByTime( gameState.TargetTime );
			gameState.TargetMultiplicator 	= defs.Float2Number( mplx );
		}
	}else
	if( gameState.GameState == ECrashGameState_WaitEnd){

		if( gameState.StateTimer > gameState.TargetTime){
			
			This.setState( serverData, ECrashGameState_Announce );
		}
	}
	
}

func (This*CrashGameLogic) InitMainThread(serverData* DataServer){

	go func (){

		for iLoop := 0; ; iLoop++{
			
			time.Sleep( 10 * time.Millisecond );
			This.updateLoop( serverData );
		}
	}()
}