package models

type ECrashGameState string

const (
	ECrashGameState_Announce = "announce"
	ECrashGameState_Starting = "starting"
	ECrashGameState_Start    = "start"
	ECrashGameState_Playing  = "playing"
	ECrashGameState_WaitEnd  = "wait"
)