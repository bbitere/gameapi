package models

type Player struct {
	PlayerName                     string
	IsHidden                       bool
	CashoutMultiplicator           int //defs.TDecimal
	NextRound_CashoutMultiplicator int //defs.TDecimal
}