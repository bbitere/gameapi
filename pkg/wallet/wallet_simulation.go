package wallet

type Mokup_Wallet struct {
	wallet TMoney
}

func (This *Mokup_Wallet) InitWallet() {

	This.wallet = 0
}

func (This *Mokup_Wallet) GetWallet() TMoney {
	return This.wallet
}
