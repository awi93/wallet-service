package wallet

import (
	walletSvc "github.com/awi93/wallet-service/src/services/wallet"
)

type Deposit struct {
	service walletSvc.Service
}

type Args struct {
	Service walletSvc.Service
}

func NewHandler(Args *Args) *Deposit {
	return &Deposit{
		service: Args.Service,
	}
}
