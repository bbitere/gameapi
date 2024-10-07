package models

import (
	time "time"
)

type Player struct {

	Token                          	string
	PlayerName                     	string
	IsHidden                       	bool
	Currency						string
	CashedoutAtValue				int32//if player cashed out the money at a certain value.
	TimeStamp						time.Time;
	CashoutMultiplicator           	int32 //defs.TDecimal
	NextRound_CashoutMultiplicator 	int32 //defs.TDecimal
}