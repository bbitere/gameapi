package models

type GameState struct {
	GameState  string //ECrashGameState
	StateTimer float64

	TargetMultiplicator int32   //defs.TDecimal
	TargetTime          float64 //defs.TDecimal
}