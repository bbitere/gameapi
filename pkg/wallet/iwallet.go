package wallet

type TMoney = int32
type IWallet interface {
	GetWallet() TMoney
	InitWallet()
}