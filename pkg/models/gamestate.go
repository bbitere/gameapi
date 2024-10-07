package models


type GameState struct{

	GameState				ECrashGameState		
	GameWillStartInSecond	float64
	TargetMultiplicator		int32 //defs.TDecimal
}